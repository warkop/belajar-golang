package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type student struct {
	id    string
	name  string
	age   int
	grade int
}

func connect() (*sql.DB, error) {
	// db, err := sql.Open("mysql", "root:toor@tcp(127.0.0.1:3306)/belajar_golang") ini yang mysql
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=belajar_golang sslmode=disable")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func sqlExec() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("insert into student values ($1, $2, $3, $4)", "G001", "Galahad", 29, 2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("insert success!")

	_, err = db.Exec("update student set age = $1 where id = $2", 28, "G001")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("update success!")

	_, err = db.Exec("delete from student where id = $1", "G001")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("delete success!")

	//jika menggunakan prepand
	// menggunakan metode prepared statement
	// stmt, err := db.Prepare("insert into tb_student values (?, ?, ?, ?)")
	// stmt.Exec("G001", "Galahad", 29, 2)
}

func main() {
	sqlExec()
}
