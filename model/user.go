package model

import (
	"fmt"

	"github.com/google/uuid"
	. "github.com/usermanage/db"
	"github.com/usermanage/util"
)

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Account  string `json:"account"`
	Password string `json:password`
}

func SaveUser(u *User) {
	stmt, err := Db.Prepare("INSERT sys_user SET id=?, name=?, account=?,Password=?")
	checkErr(err)
	id := uuid.New().String()
	_, err = stmt.Exec(id, u.Name, u.Account, u.Password)
	checkErr(err)
	u.Id = id
}

func UserList(pager util.Pager) util.PageResult {
	pr := util.PageResult{}
	sql := "SELECT * FROM sys_user"

	countStmt, err := Db.Prepare("SELECT count(*) FROM (" + sql + ") a")
	checkErr(err)
	countStmt.QueryRow().Scan(&pr.Total)

	rowsStmt, err := Db.Prepare(sql + fmt.Sprintf(" LIMIT %v, %v", pager.Start(), pager.Limit()))
	checkErr(err)

	rows, err := rowsStmt.Query()
	checkErr(err)

	records := []User{}
	for rows.Next() {
		u := User{}
		err = rows.Scan(&u.Id, &u.Name, &u.Account, &u.Password)
		checkErr(err)
		records = append(records, u)
	}
	pr.Records = records
	return pr
}

func GetUser(account, password string) *User {
	stmt, err := Db.Prepare("SELECT * FROM sys_user WHERE account=? AND password=?")
	checkErr(err)
	rows, err := stmt.Query(account, password)
	checkErr(err)
	for rows.Next() {
		u := User{}
		err = rows.Scan(&u.Id, &u.Name, &u.Account, &u.Password)
		checkErr(err)
		return &u
	}
	return nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
