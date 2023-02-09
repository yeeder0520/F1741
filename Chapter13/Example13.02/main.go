package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

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

	DBCreate := `
	CREATE TABLE employee
	(
		id INT NOT NULL UNIQUE,
		name VARCHAR(20)
	);`

	_, err = db.Exec(DBCreate)
	if err != nil {
		panic(err)
	}
	fmt.Println("表格 employee 已建立")
}
