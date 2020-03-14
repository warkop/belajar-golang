package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type student struct {
	id    string
	name  string
	age   int
	grade int
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:toor@tcp(127.0.0.1:3306)/belajar_golang")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func sqlQueryRow() {
	var db, err = connect()
	if err != nil {
		fmt.Println(err.Error())
	}

	defer db.Close()

	var result = student{}
	var id = "E001"
	err = db.
		QueryRow("select name, grade from student where id = ?", id).
		Scan(&result.name, &result.grade)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("name: %s\ngrade: %d\n", result.name, result.grade)
}

func main() {
	sqlQueryRow()
}
