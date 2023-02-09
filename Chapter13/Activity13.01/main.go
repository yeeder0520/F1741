package main

import (
	"database/sql"
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/ssh/terminal"
)

const (
	id = iota
	payee
	spent
	category
)

const (
	dbType = "mysql"
	dbName = "mysqldb"
	dbHost = "localhost"
	dbPort = 3306
)

var budgetCategory = map[string]string{
	"fuel":           "autoFuel",
	"gas":            "autoFuel",
	"mortgage":       "mortgage",
	"repairs":        "repairs",
	"car insurance":  "insurance",
	"life insurance": "insurance",
	"utilities":      "utilities",
}

var (
	ErrOpenCSV               = errors.New("CSV 檔開啟錯誤")
	ErrReadCSVRow            = errors.New("CSV 檔單行讀取錯誤")
	ErrInvalidID             = errors.New("ID 錯誤")
	ErrInvalidPayee          = errors.New("收款人錯誤")
	ErrInvalidSpent          = errors.New("金額錯誤")
	ErrInvalidBudgetCategory = errors.New("查無分類")
	ErrDBConnect             = errors.New("無法連接資料庫")
	ErrDBCreateTable         = errors.New("建立資料表失敗")
	ErrDBInsert              = errors.New("於資料表插入資料失敗")
	ErrDBQuery               = errors.New("查詢資料表失敗")
	ErrDBReadRow             = errors.New("讀取單筆資料失敗")
)

type transaction struct {
	id       int
	payee    string
	spent    float64
	category string
}

func main() {
	bankFileName := flag.String("source", "", "銀行交易檔 (CSV) 路徑與檔名")
	logFileName := flag.String("log", "", "日誌檔路徑與檔名")

	flag.Parse()
	if *bankFileName == "" {
		fmt.Println("需要指定銀行交易檔 (CSV)")
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *logFileName == "" {
		fmt.Println("需要指定日誌檔")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if _, err := os.Stat(*bankFileName); errors.Is(err, os.ErrNotExist) {
		fmt.Println("查無銀行交易檔 (CSV):", *bankFileName)
		os.Exit(1)
	}

	var userName string
	fmt.Print("請輸入資料庫帳號: ")
	fmt.Scanln(&userName)
	if len(userName) == 0 {
		fmt.Println("未輸入資料庫帳號")
		os.Exit(1)
	}

	fmt.Print("請輸入資料庫密碼: ")
	password, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Println("未輸入資料庫密碼")
		os.Exit(1)
	}
	fmt.Println("")

	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)

	trs := parseBankFile(*bankFileName, *logFileName)
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", userName, string(password), dbHost, dbPort, dbName)
	tableName := strings.TrimSuffix(*bankFileName, filepath.Ext(*bankFileName))

	writeDatabase(dataSource, tableName, trs, *logFileName)
}

func writeDatabase(dataSource, tableName string, trs []transaction, logFileName string) {
	db, _ := sql.Open(dbType, dataSource)
	defer db.Close()
	if err := db.Ping(); err != nil {
		writeLog(logFileName, true, ErrDBConnect, err.Error())
	}

	DBDrop := "DROP TABLE " + tableName + ";"
	db.Exec(DBDrop)

	DBCreate := `
	CREATE TABLE ` + tableName + `
	(
		id INT NOT NULL UNIQUE,
		payee VARCHAR(20),
		spent FLOAT,
		category VARCHAR(20)
	);`

	if _, err := db.Exec(DBCreate); err != nil {
		writeLog(logFileName, true, ErrDBCreateTable, err.Error()+": "+tableName)
	}

	DBInsert := "INSERT INTO " + tableName + " (id, payee, spent, category) VALUES (?, ?, ?, ?);"
	insertStmt, err := db.Prepare(DBInsert)
	for err != nil {
		writeLog(logFileName, true, err, "")
	}
	for _, tr := range trs {
		if tr.id == 0 || tr.payee == "" || tr.spent == 0 || tr.category == "" {
			continue
		}
		if _, err := insertStmt.Exec(tr.id, tr.payee, tr.spent, tr.category); err != nil {
			writeLog(logFileName, false, ErrDBInsert, fmt.Sprintf("%v: id=%v", err.Error(), tr.id))
		}
	}
	insertStmt.Close()

	DBQuery := "SELECT * FROM " + tableName + ";"
	rows, err := db.Query(DBQuery)
	if err != nil {
		writeLog(logFileName, false, ErrDBQuery, "")
	}

	fmt.Println("DB 查詢結果:")
	for rows.Next() {
		tr := transaction{}
		if err := rows.Scan(&tr.id, &tr.payee, &tr.spent, &tr.category); err != nil {
			writeLog(logFileName, false, ErrDBReadRow, fmt.Sprintf("%v: id=%v", err.Error(), tr.id))
		}
		fmt.Printf("Id: %-2v\tPayee: %-15v\tSpent: %-5v\tCateory: %-10v\n", tr.id, tr.payee, tr.spent, tr.category)
	}
	if err := rows.Err(); err != nil {
		writeLog(logFileName, true, err, "")
	}
	rows.Close()

}

func parseBankFile(bankFileName, logFileName string) []transaction {
	bankFile, err := os.Open(bankFileName)
	if err != nil {
		writeLog(logFileName, true, ErrOpenCSV, bankFileName)
	}

	r := csv.NewReader(bankFile)
	header := true
	trs := []transaction{}

	for {
		tr := transaction{}
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			writeLog(logFileName, false, ErrReadCSVRow, err.Error())
		}

		if header {
			header = false
			continue
		}
		for index, value := range record {
			switch index {
			case id:
				id, err := strconv.Atoi(strings.TrimSpace(value))
				if err != nil {
					writeLog(logFileName, false, ErrInvalidID, "Index: "+fmt.Sprint(index), value)
				}
				tr.id = id
			case payee:
				value = strings.TrimSpace(value)
				if len(value) == 0 {
					writeLog(logFileName, false, ErrInvalidPayee, "Id: "+fmt.Sprint(tr.id), value)
				}
				tr.payee = value
			case spent:
				value = strings.TrimSpace(value)
				spent, err := strconv.ParseFloat(value, 64)
				if err != nil {
					writeLog(logFileName, false, ErrInvalidSpent, "Id: "+fmt.Sprint(tr.id), value)
				}
				tr.spent = spent
			case category:
				data, ok := budgetCategory[strings.TrimSpace(value)]
				if !ok {
					writeLog(logFileName, false, ErrInvalidBudgetCategory, "Id: "+fmt.Sprint(tr.id), value)
				}
				tr.category = data
			}
		}
		trs = append(trs, tr)
	}
	return trs

}

func writeLog(logFileName string, fatal bool, errMsg error, data ...string) {
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Panicln(err)
	}
	defer logFile.Close()

	logger := log.New(logFile, "", log.Ldate|log.Lmicroseconds|log.Llongfile)
	logger.Println(errMsg, data)

	if !fatal {
		log.Println(errMsg, data)
	} else {
		log.Panicln(errMsg, data)
	}
}
