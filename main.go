package main

import (
	"github.com/usermanage/db"
	"github.com/usermanage/route"
	"github.com/usermanage/util"
)

func init() {
	util.Parse()
}

func main() {
	router := route.InitRouter()
	db.InitMysqlDb()
	router.Run(":8888")
}
