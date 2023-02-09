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
	if err := db.Ping(); err != nil {
		panic(err)
	}

	DBDrop := "DROP TABLE fizzbuzz;"
	db.Exec(DBDrop)

	DBCreate := `
	CREATE TABLE fizzbuzz
	(
		number INT NOT NULL,
		status VARCHAR(10)
	);`
	if _, err := db.Exec(DBCreate); err != nil {
		panic(err)
	}

	DBInsert := "INSERT INTO fizzbuzz (number, status) VALUES (?, ?);"
	insertStmt, err := db.Prepare(DBInsert)
	for err != nil {
		panic(err)
	}
	for i := 1; i <= 100; i++ {
		if _, err := insertStmt.Exec(i, ""); err != nil {
			fmt.Println(err)
		}
	}
	insertStmt.Close()
	fmt.Println("插入資料筆數: 100")

	DBUpdate := "UPDATE fizzbuzz SET status=? WHERE MOD(number, ?)=0 AND status=''"
	updateStmt, err := db.Prepare(DBUpdate)
	if err != nil {
		panic(err)
	}
	numbers := []int{15, 3, 5}
	statuses := []string{"FizzBuzz", "Fizz", "Buzz"}
	for i := 0; i < 3; i++ {
		updatedResult, err := updateStmt.Exec(statuses[i], numbers[i])
		if err != nil {
			panic(err)
		}
		updatedRecords, err := updatedResult.RowsAffected()
		if err != nil {
			panic(err)
		}
		fmt.Println(statuses[i], "更新筆數:", updatedRecords)
	}
	updateStmt.Close()

	DBDelete := "DELETE FROM fizzbuzz WHERE status=?"
	deleteStmt, err := db.Prepare(DBDelete)
	if err != nil {
		panic(err)
	}
	deletedResult, err := deleteStmt.Exec("FizzBuzz")
	if err != nil {
		panic(err)
	}
	deletedRecords, err := deletedResult.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("FizzBuzz 刪除筆數:", deletedRecords)

	rowStmt, err := db.Prepare("SELECT COUNT(*) FROM fizzbuzz")
	if err != nil {
		panic(err)
	}
	var count int
	if err := rowStmt.QueryRow().Scan(&count); err != nil {
		panic(err)
	}
	fmt.Println("資料表總筆數: ", count)
	rowStmt.Close()
}
