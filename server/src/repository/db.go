package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ssizawa/taskapp/server/src/structure"
)

var Db *sql.DB

func Opendb() {
	Db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/taskapp?")
	if err != nil {
		fmt.Print("========1========")
		panic(err.Error())
	}
	defer Db.Close()
}

func VerifybyName(name string, password string) bool {
	var user structure.User
	_ = Db.QueryRow("SELECT name FROM User WHERE name = ? AND password = ?", name, password).Scan(&user.Name)
	if user.Name == "" {
		return false
	}
	return true
}
