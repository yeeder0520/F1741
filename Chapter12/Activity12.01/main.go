package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	id = iota
	payee
	spent
	category
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

	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)

	trs := parseBankFile(*bankFileName, *logFileName)
	fmt.Println("CSV 檔解析結果:")
	for _, tr := range trs {
		fmt.Printf("Id: %-2v\tPayee: %-15v\tSpent: %-5v\tCateory: %-10v\n", tr.id, tr.payee, tr.spent, tr.category)
	}
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
