package users

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type User struct {
	db  *sqlx.DB
	log zerolog.Logger
}

func New(db *sqlx.DB, log zerolog.Logger) User {
	return User{
		db:  db,
		log: log,
	}
}

func (u User) Create(ctx context.Context, info Info) (Info, error) {
	info.IsActive = true
	q := `INSERT INTO public.user (user_email,first_name,last_name,is_active)
	       values(:user_email,:first_name,:last_name,:is_active) 
		   returning id`
	stmt, err := u.db.PrepareNamed(q)
	if err != nil {
		return info, errors.Wrap(err, "Create User")
	}
	err = stmt.Get(&info.ID, info)
	if err != nil {
		return info, errors.Wrap(err, "Get record")
	}

	return info, nil
}

func (u User) Get(ctx context.Context, userID uint64) (Info, error) {
	info := Info{}
	q := `SELECT id,user_email,first_name,last_name,is_active FROM public.user
		WHERE id=$1`
	err := u.db.GetContext(ctx, &info, q, userID)
	if err != nil && err != sql.ErrNoRows {
		return info, errors.Wrap(err, "Get User")
	}

	return info, nil
}

func (u User) Update(ctx context.Context, info Info) error {
	q := "UPDATE public.user SET first_name=$2,last_name=$3 WHERE id=$1"
	_, err := u.db.ExecContext(ctx, q, info.ID, info.FirstName, info.LastName)
	if err != nil {
		return errors.Wrap(err, "Update User")
	}
	return nil
}

func (u User) Delete(ctx context.Context, userID uint64) error {
	q := "DELETE FROM public.user WHERE id=$1"
	_, err := u.db.ExecContext(ctx, q, userID)
	if err != nil {
		return errors.Wrap(err, "Deleteing User")
	}

	return nil
}

func (u User) InActived(ctx context.Context, userIDs []uint64) error {
	q := "UPDATE public.user SET is_active=false WHERE id =$1"

	for _, key := range userIDs {
		_, err := u.db.ExecContext(ctx, q, key)
		if err != nil {
			return errors.Wrap(err, "Update User In-Actived")
		}
	}

	return nil
}

func (u User) GetAll(ctx context.Context) ([]Info, error) {
	info := []Info{}
	q := `SELECT id,user_email,first_name,last_name,is_active FROM public.user`
	err := u.db.Select(&info, q)
	if err != nil && err != sql.ErrNoRows {
		return info, errors.Wrap(err, "Get ALL User")
	}

	return info, nil
}

func (u User) GetActiveUser(ctx context.Context) ([]Info, error) {
	info := []Info{}
	q := `SELECT id,user_email,first_name,last_name,is_active FROM public.user where is_active=true`
	err := u.db.Select(&info, q)
	if err != nil && err != sql.ErrNoRows {
		return info, errors.Wrap(err, "Get ALL User")
	}

	return info, nil
}
