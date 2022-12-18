package users

type Info struct {
	ID        uint64 `db:"id" json:"id" param:"id"`
	UserEmail string `db:"user_email" json:"user_email"`
	FirstName string `db:"first_name" json:"first_name"`
	LastName  string `db:"last_name" json:"last_name"`
	IsActive  bool   `db:"is_active" json:"is_active"`
}
