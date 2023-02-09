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

	updateStmt, err := db.Prepare("UPDATE employee SET name=? WHERE id=?")
	if err != nil {
		panic(err)
	}
	defer updateStmt.Close()

	e := employee{204, "Robert"}
	updatedResult, err := updateStmt.Exec(e.name, e.id)
	if err != nil {
		panic(err)
	}
	updatedRecords, err := updatedResult.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("更新資料筆數:", updatedRecords)
}
