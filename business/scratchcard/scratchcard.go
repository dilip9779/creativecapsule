package scratchcard

import (
	"context"
	"creativecapsule/business/users"
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type ScratchCard struct {
	db   *sqlx.DB
	log  zerolog.Logger
	user users.User
}

func New(db *sqlx.DB, log zerolog.Logger, user users.User) ScratchCard {
	return ScratchCard{
		db:   db,
		log:  log,
		user: user,
	}
}

func (c ScratchCard) Create(ctx context.Context, card CreateCard) ([]Info, error) {
	info := []Info{}
	data, err := c.GetActiveCards(ctx, Info{
		IsScratched: false,
		IsActive:    true,
	})
	if err != nil {
		return info, errors.Wrap(err, "getActiveCard")
	}

	if len(data) < int(card.Crad) {
		now := time.Now()
		expiryDate := now.AddDate(0, 0, 5)
		for i := 0; i < int(card.Crad)-len(data); i++ {
			discountAmount := rand.Intn(1000)
			infoSlice := Info{
				DiscountAmount: float64(discountAmount),
				ExpiryDate:     expiryDate,
			}
			fmt.Println(infoSlice)
			q := `INSERT INTO public.scratch_card (discount_amount,expiry_date)
					values(:discount_amount,:expiry_date) 
					returning id`
			stmt, err := c.db.PrepareNamed(q)
			if err != nil {
				return info, errors.Wrap(err, "Create Scratch Card")
			}
			err = stmt.Get(&infoSlice.ID, infoSlice)
			if err != nil {
				return info, errors.Wrap(err, "Get record")
			}
			info = append(info, infoSlice)
		}
	} else {
		return info, errors.New(fmt.Sprintf("%d number of active scratch cards still exists in the DB. Did not create any new scratch card", len(data)))
	}

	return info, nil
}

func (c ScratchCard) Get(ctx context.Context) ([]Info, error) {
	info := []Info{}
	q := `SELECT id,discount_amount,expiry_date,is_scratched,is_active FROM public.scratch_card`
	err := c.db.Select(&info, q)
	if err != nil {
		return info, errors.Wrap(err, "Get Scratch Card")
	}

	return info, nil
}

func (c ScratchCard) Update(ctx context.Context, info Info) error {
	q := "UPDATE public.scratch_card SET is_scratched=true WHERE id=$1"
	_, err := c.db.ExecContext(ctx, q, info.ID)
	if err != nil {
		return errors.Wrap(err, "Update Is scartched")
	}

	return nil
}

func (c ScratchCard) DeActived(ctx context.Context, info Info) error {
	q := "UPDATE public.scratch_card SET is_active=false WHERE id=$1"
	_, err := c.db.ExecContext(ctx, q, info.ID)
	if err != nil {
		return errors.Wrap(err, "Update Is active")
	}

	return nil
}

func (c ScratchCard) GetActiveCards(ctx context.Context, filter Info) ([]Info, error) {
	info := []Info{}
	q := `SELECT id,discount_amount,expiry_date,is_scratched,is_active FROM public.scratch_card
		WHERE is_scratched=:is_scratched and is_active=:is_active and expiry_date >= CAST(NOW() as date)`
	if filter.ID > 0 {
		q += ` AND id =:id`
	}
	c.log.Info().Msgf("Filter %+v Query :%s", filter, q)
	stmt, err := c.db.PrepareNamedContext(ctx, q)
	if err != nil {
		return info, errors.Wrap(err, "Stmt Scratch Card")
	}
	err = stmt.Select(&info, filter)
	if err != nil && err != sql.ErrNoRows {
		return info, errors.Wrap(err, "Get Scratch Card")
	}

	return info, nil
}

func (c ScratchCard) TransactionUpdate(ctx context.Context, filter Filter) error {
	scratchCard, err := c.GetActiveCards(ctx, Info{
		ID:          filter.ScratchCardID,
		IsScratched: false,
		IsActive:    true,
	})
	if err != nil {
		return errors.Wrap(err, "Get Scratch Card")
	}

	if len(scratchCard) == 0 {
		return errors.New("Scratch Card Not Valid")
	}

	userInfo, err := c.user.Get(ctx, filter.UserID)
	if err != nil {
		return errors.Wrap(err, "Get User Info")
	}
	if userInfo.ID == 0 {
		return errors.New("User Not Valid")
	}
	filter.TransactionAmount = scratchCard[0].DiscountAmount
	query := `INSERT INTO public.transaction ( transaction_amount, user_id, scratch_card_id)
					values(:transaction_amount,:user_id,:scratch_card_id) `

	_, err = c.db.NamedExec(query, filter)
	c.log.Debug().Msgf("%+v", err)
	if err != nil {
		return errors.Wrap(err, "Insert transaction Card")
	}

	err = c.Update(ctx, Info{
		ID: filter.ScratchCardID,
	})
	if err != nil {
		return errors.Wrap(err, "Get Scratch Card")
	}

	return nil
}

func (c ScratchCard) GetTransaction(ctx context.Context, filter Filter) ([]TransactionInfo, error) {
	info := []TransactionInfo{}
	q := `SELECT 
				t.transaction_amount as amount,
				t.date_of_transaction,
				u.first_name as user_name
			FROM public.transaction t
			LEFT JOIN public.user u ON u.id = t.user_id
			WHERE 1=1`
	if filter.UserID > 0 {
		q += ` AND u.id =:user_id`
	}
	if filter.TransactionAmount > 0 {
		q += ` AND t.transaction_amount =:transaction_amount`
	}

	if len(filter.TransactionDate) > 0 {
		q += ` AND to_char(t.date_of_transaction,'YYYY-MM-DD') =:date_of_transaction`
	}
	c.log.Debug().Msg(q)
	stmt, err := c.db.PrepareNamedContext(ctx, q)
	if err != nil {
		return info, errors.Wrap(err, "Stmt Scratch Card")
	}
	err = stmt.Select(&info, filter)
	if err != nil && err != sql.ErrNoRows {
		return info, errors.Wrap(err, "Get Scratch Card")
	}

	return info, nil
}
