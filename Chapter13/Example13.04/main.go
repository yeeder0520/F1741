package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type employee struct {
	id   int
	name string
}

func main() {
	db, err := sql.Open("mysql", "user:1234@tcp(localhost:3306)/mysqldb?charset=utf8")

	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("sql.DB 結構已建立")

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("資料庫連線成功")

	rows, err := db.Query("SELECT * FROM employee")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	fmt.Println("資料表查詢成功, 列出 employee 內容...")

	for rows.Next() {
		e := employee{}
		err := rows.Scan(&e.id, &e.name)
		if err != nil {
			panic(err)
		}
		fmt.Println(e.id, e.name)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
}
