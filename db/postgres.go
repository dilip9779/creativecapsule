package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// GetSQLXPGInstance returns pointer to sqlx.DB
func GetSQLXPGInstance(
	username string,
	password string,
	host string,
	port string,
	dbName string, sslMode string) (*sqlx.DB, error) {

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host,
		port,
		username,
		password,
		dbName,
		sslMode,
	)

	return sqlx.Connect("postgres", dsn)
}

// GetSQLXPGInstance returns pointer to sqlx.DB
func GetSQLXPGInstanceFromDsn(dsn string) (*sqlx.DB, error) {
	return sqlx.Connect("postgres", dsn)
}
