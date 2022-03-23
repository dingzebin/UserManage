package main

import (
	"github.com/usermanage/db"
	"github.com/usermanage/route"
)

func main() {
	router := route.InitRouter()
	db.InitMysqlDb()
	router.Run(":8888")
}
