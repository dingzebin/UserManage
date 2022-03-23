package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var Db *sql.DB

func InitMysqlDb() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/user_manage")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	Db = db
}
