package infrastructure

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	TeamTable         = "team"
	UserTable         = "user"
	SubscriptionTable = "subscription"
	NotificationTable = "notification"
)

func MakeMySql() *sql.DB {
	port := os.Getenv("APP_DB_PORT")
	if port == "" {
		port = "3306"
	}

	connection, err := sql.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&multiStatements=true",
		os.Getenv("APP_DB_USER"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_HOST"),
		port,
		os.Getenv("APP_DB_NAME"),
	))

	if err != nil {
		panic(err.Error())
	}

	return connection
}
