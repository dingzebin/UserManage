package model

import (
	"github.com/google/uuid"
	. "github.com/usermanage/db"
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

func UserList() {

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
