package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// GetSQLXMysqlInstance returns pointer to sqlx.DB
func GetSQLXMysqlInstance(
	username string,
	password string,
	host string,
	port string,
	dbName string) (*sqlx.DB, error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		username,
		password,
		host,
		port,
		dbName,
	)

	return sqlx.Connect("mysql", dsn)
}

// GetSQLXMysqlInstance returns pointer to sqlx.DB from dns
func GetSQLXMysqlInstanceFromDns(dsn string) (*sqlx.DB, error) {
	return sqlx.Connect("mysql", dsn)
}
