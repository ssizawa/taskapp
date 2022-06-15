package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Create_table() {
	db, err := sql.Open("mysql", "roor:root@tcp(127.0.0.1:3306)/taskapp")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM User")
	if err != nil {
		panic(err.Error())
	}
	fmt.Print(rows)

	// columns, err := rows.Columns()
	// if err != nil {
	// 	panic(err.Error())
	// }

	// values := make([]sql.RawBytes, len(columns))

	// scanArgs := make([]interface{}, len(values))
	// for i := range values {
	// 	scanArgs[i] = &values[i]
	// }

	// for rows.Next() {
	// 	err = rows.Scan(scanArgs...)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}

	// 	var value string
	// 	for i, col := range values {
	// 		if col == nil {
	// 			value = "NULL"
	// 		} else {
	// 			value = string(col)
	// 		}
	// 		fmt.Println(columns[i], ":", value)
	// 	}
	// 	fmt.Println("-----------------------------------")
	// }
}

// func InsertUser(name, pw) {
// 	ins, err := db.Prepare("INSET INTO User(name, password) VALUES(?, ?)")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer ins.Close()

// 	result, err := ins.Exec(name, pw)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// }
