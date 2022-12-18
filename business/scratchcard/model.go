package scratchcard

import "time"

type Info struct {
	ID             uint64    `db:"id" json:"id" param:"id"`
	DiscountAmount float64   `db:"discount_amount" json:"discount_amount"`
	ExpiryDate     time.Time `db:"expiry_date" json:"expiry_date"`
	IsScratched    bool      `db:"is_scratched" json:"is_scratched"`
	IsActive       bool      `db:"is_active" json:"is_active"`
}

type CreateCard struct {
	Crad uint64 `json:"card"`
}

type Filter struct {
	UserID            uint64  `db:"user_id" json:"user_id"`
	ScratchCardID     uint64  `db:"scratch_card_id" json:"scratch_card_id"`
	TransactionAmount float64 `db:"transaction_amount" json:"amount"`
	TransactionDate   string  `db:"date_of_transaction" json:"date_of_transaction"`
}

type TransactionInfo struct {
	Amount          float64 `db:"amount" json:"amount"`
	TransactionDate string  `db:"date_of_transaction" json:"date_of_transaction"`
	UserName        string  `db:"user_name" json:"user_name"`
}
