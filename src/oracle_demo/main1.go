package main

import (
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-oci8"
)

func main() {
	const (
		openString = "system/oracle@localhost:1521"
		tableName  = "DOCUMENT_KB"
	)

	db, err := sql.Open("oci8", openString)
	if err != nil {
		fmt.Printf("Open error is not nil: %v", err)
		return
	}

	if db == nil {
		fmt.Println("db is nil")
		return
	}

	// defer close database
	defer func() {
		err = db.Close()
		if err != nil {
			fmt.Println("Close error is not nil:", err)
		}
	}()

	query := "select * from " + tableName

	var rows *sql.Rows
	ctx, cancel := context.WithTimeout(context.Background(), 55*time.Second)
	defer cancel()
	if rows, err = db.QueryContext(ctx, query); err != nil {
		log.Println("QueryContext: ", err)
		return
	}
	defer rows.Close()

	var (
		documentNo   int64
		barCodeNo    string
		businessType string
		documentType string
		documentCode string
		scanTime     string //time
		scanSite     string
		scanOperator string
		pageCount    int64
		regionCode   string
		status       string
		creationDate string      //time
		lcd          interface{} //time
		lcu          interface{}
		documentId   interface{}
		fileType     string
		documentSize int64
		priority     interface{}
		scanSeqNo    interface{}
		isMigrated   string
	)

	if columns, err := rows.Columns(); err != nil {
		log.Println("err ========", err)
		return
	} else {
		log.Println("=======columns: ", columns)
	}

	for rows.Next() {
		if err = rows.Scan(&documentNo, &barCodeNo, &businessType, &documentType, &documentCode,
			&scanTime, &scanSite, &scanOperator, &pageCount, &regionCode, &status, &creationDate,
			&lcd, &lcu, &documentId, &fileType, &documentSize, &priority, &scanSeqNo, &isMigrated,
		); err != nil {
			fmt.Println("Scan error is not nil:", err)
			return
		}

		fmt.Println("***********************************************")
		fmt.Println("===documentNo: ", documentNo)
		fmt.Println("===barCodeNo: ", barCodeNo)
		fmt.Println("===businessType: ", businessType)
		fmt.Println("===documentType: ", documentType)
		fmt.Println("===documentCode: ", documentCode)
		fmt.Println("===scanTime: ", scanTime)
		fmt.Println("===scanSite: ", scanSite)
		fmt.Println("===scanOperator: ", scanOperator)
		fmt.Println("===regionCode: ", regionCode)
		fmt.Println("===status: ", status)
		fmt.Println("===creationDate: ", creationDate)
		fmt.Println("===lcd: ", lcd)
		fmt.Println("===lcu: ", lcu)
		fmt.Println("===documentId: ", documentId)
		fmt.Println("===fileType: ", fileType)
		fmt.Println("===documentSize: ", documentSize)
		fmt.Println("===priority: ", priority)
		fmt.Println("===scanSeqNo: ", scanSeqNo)
		fmt.Println("===isMigrated: ", isMigrated)
		fmt.Println("***********************************************")
		fmt.Printf("\n")

		if _, err := db.ExecContext(ctx, "update document_kb set IS_MIGRATED = 'Y' where document_No = :1", documentNo); err != nil {
			log.Println("========== err: ", err)
		}
	}

	err = rows.Err()
	if err != nil {
		fmt.Println("Err error is not nil:", err)
		return
	}
}

func CalcSignatureForBinary(data []byte, key string) []byte {
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write(data)
	return mac.Sum(nil)
}

/// update document_kb set IS_MIGRATED = 'Y' where DOCUMENT_NO = '890307';
