package main

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-oci8"
)

type KBDocument struct {
	DocumentNo   *int64  `db:"DOCUMENT_NO"`
	BarCodeNo    *string `db:"BAR_CODE_NO"`
	BusinessType *string `db:"BUSINESS_TYPE"`
	DocumentType *string `db:"DOCUMENT_TYPE"`
	DocumentCode *string `db:"DOCUMENT_CODE"`
	ScanTime     *string `db:"SCAN_TIME"`
	ScanSite     *string `db:"SCAN_SITE"`
	ScanOperator *string `db:"SCAN_OPERATOR"`
	PageCount    *int32  `db:"PAGE_COUNT"`
	RegionCode   *string `db:"REGION_CODE"`
	Status       *string `db:"STATUS"`
	CreationDate *string `db:"CREATION_DATE"`
	Lcd          *string `db:"LCD"`
	Lcu          *string `db:"LCU"`
	DocumentID   *int32  `db:"DOCUMENT_ID"`
	FileType     *string `db:"FILE_TYPE"`
	DocumentSize *int64  `db:"DOCUMENT_SIZE"`
	Priority     *int8   `db:"PRIORITY"`
	ScanSeqNo    *string `db:"SCAN_SEQNO"`
	IsMigrated   *string `db:"IS_MIGRATED"`
}

type BizdataValue interface{}
type Bizdata map[string]BizdataValue

type Document struct {
	DocId                   string    `bson:"doc_id"`
	AppId                   string    `bson:"app_id"`
	Title                   string    `bson:"title"`
	Bizdata                 Bizdata   `bson:"bizdata"`
	Status                  string    `bson:"status,omitempty"`
	DecompressionFailedKeys []string  `bson:"decompression_failed_keys,omitempty"`
	ExpiredAt               time.Time `bson:"expired_at,omitempty"`
}

func main() {
	var (
		rows *sqlx.Rows
		doc  KBDocument
		ctx  context.Context
		db   *sqlx.DB
		err  error
	)
	const (
		openString = "system/oracle@localhost:1521"
		tableName  = "document_kb"
		limit      = 3
	)

	if db, err = sqlx.Connect("oci8", openString); err != nil {
		log.Fatalf("Failed to connect db: %s", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if rows, err = db.QueryxContext(ctx, "select * from document_kb where rownum <= :1", limit); err != nil {
		log.Println("QueryContext: ", err)
		return
	}
	defer func() {
		rows.Close()
		fmt.Println("======= close")
	}()

	var docIds []string
	for rows.Next() {
		if err = rows.StructScan(&doc); err != nil {
			log.Fatalf("Failed to scan doc: %s", err)
		}

		fmt.Println("***********************************************")
		fmt.Println("===documentNo: ", *doc.DocumentNo)
		// fmt.Println("===barCodeNo: ", *doc.BarCodeNo)
		// fmt.Println("===businessType: ", *doc.BusinessType)
		// fmt.Println("===documentType: ", *doc.DocumentType)
		// fmt.Println("===documentCode: ", *doc.DocumentCode)
		// fmt.Println("===scanTime: ", *doc.ScanTime)
		// fmt.Println("===scanSite: ", *doc.ScanSite)
		// fmt.Println("===scanOperator: ", *doc.ScanOperator)
		// fmt.Println("===regionCode: ", *doc.RegionCode)
		// fmt.Println("===status: ", *doc.Status)
		// fmt.Println("===creationDate: ", doc.CreationDate)
		// fmt.Println("===lcd: ", doc.Lcd)
		// fmt.Println("===lcu: ", doc.Lcu)
		// fmt.Println("===documentId: ", doc.DocumentID)
		// fmt.Println("===fileType: ", doc.FileType)
		// fmt.Println("===documentSize: ", doc.DocumentSize)
		// fmt.Println("===priority: ", doc.Priority)
		// fmt.Println("===scanSeqNo: ", doc.ScanSeqNo)
		fmt.Println("===isMigrated: ", *doc.IsMigrated)
		fmt.Println("***********************************************")
		fmt.Printf("\n")

		var d = Document{
			Bizdata: make(Bizdata),
		}
		v := reflect.ValueOf(doc)
		st := reflect.TypeOf(doc)

		for i := 0; i < st.NumField(); i++ {
			field := st.Field(i)
			bizValue := v.Field(i)
			if bizValue.IsNil() {
				continue
			}
			bizKey := field.Tag.Get("db")

			d.Bizdata[bizKey] = bizValue.Elem()
		}

		for k, v := range d.Bizdata {
			fmt.Printf("===k: %s v: %v\n", k, v)
		}

		fmt.Printf("========== %#v\n", d)

		docIds = append(docIds, strconv.Itoa(int(*doc.DocumentNo)))

		// docIdsString := fmt.Sprintf("(%s,%s,%s)")
		// db.MustExecContext(ctx, "update document_kb set IS_MIGRATED = 'Z' where document_No = :1", *doc.DocumentNo)
	}

	rows.Close()
	rows.Close()

	var b strings.Builder
	docIdsString := strings.Join(docIds, ",")
	b.WriteString(docIdsString)
	fmt.Println(b.String())
	m := map[string]interface{}{"IS_MIGRATED": "Z"}
	queryStmt := fmt.Sprintf("update document_kb set IS_MIGRATED=:m")

	arg := map[string]interface{}{
		"published": true,
		"authors": []{8, 19, 32, 44},
	}
	query, args, err := sqlx.Named("SELECT * FROM articles WHERE published=:published AND author_id IN (:authors)", arg)
	query, args, err := sqlx.In(query, args...)
	query = db.Rebind(query)
	db.Query(query, args...)
	// fmt.Println("====== stmt: ", queryStmt)
	// db.NamedExecContext(ctx, queryStmt, m)
	// if err = rows.Err(); err != nil {
	// 	log.Fatalf("Err error is not nil:", err)
	// }
}
