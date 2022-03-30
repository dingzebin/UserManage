package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/magiconair/properties"
	"github.com/usermanage/util"
)

var Db *sql.DB

func InitMysqlDb() {
	p := properties.MustLoadFile(util.Args.Db, properties.UTF8)
	db, err := sql.Open("mysql", p.MustGetString("url"))
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * time.Duration(p.GetInt("connMaxLifetime", 3)))
	db.SetMaxOpenConns(p.GetInt("maxOpenConns", 10))
	db.SetMaxIdleConns(p.GetInt("maxIdleConns", 10))
	Db = db
}
