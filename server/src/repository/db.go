package repository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/ssizawa/taskapp/server/src/structure"
)

var db *sql.DB

func Opendb() {

	db_name, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/taskapp?")

	if err != nil {
		panic(err.Error())
	}

	db = db_name
}

func Verify(name string, password string) bool {

	var user structure.User

	Opendb()
	defer db.Close()

	_ = db.QueryRow("SELECT name FROM User WHERE name = ? AND password = ?", name, password).Scan(&user.Name)

	if user.Name == "" {
		return false
	}

	return true
}

func ChangePass(name string, newpass string) {

	Opendb()
	defer db.Close()

	_ = db.QueryRow("UPDATE User SET password = ? WHERE name = ?", newpass, name)
}
