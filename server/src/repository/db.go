package repository

import (
	"database/sql"
	"log"

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

func VerifybyName(name string, password string) bool {

	var user structure.User

	Opendb()
	defer db.Close()

	_ = db.QueryRow("SELECT name FROM User WHERE name = ? AND password = ?", name, password).Scan(&user.Name)

	if user.Name == "" {
		return false
	}

	return true
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

func GetUserList() []string {

	var user_list []string
	var user structure.User

	Opendb()
	defer db.Close()

	rows, err := db.Query("SELECT name FROM User")

	if err != nil {
		log.Fatal(err.Error())
	}

	for rows.Next() {

		rows.Scan(&user.Name)

		user_list = append(user_list, user.Name)
	}

	return user_list
}

func CreateTask(task_name string, description string, pic string, deadline string) {

	Opendb()
	defer db.Close()

	insert, err := db.Prepare("INSERT INTO Task(task_name, description, pic, deadline, status) VALUES(?, ?, ?, ?, ?)")

	if err != nil {
		log.Fatal(err.Error())
	}

	insert.Exec(task_name, description, pic, deadline, "todo")
}
