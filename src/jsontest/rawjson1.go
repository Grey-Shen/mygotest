package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

const (
	DBEndPoint     = "localhost:27017"
	DBName         = "mydb"
	collectionName = "bizdata_test"
)

type BizdataInfo struct {
	Aa int    `bson:"aa"`
	Bb bool   `bson:"bb"`
	Cc string `bson:"cc"`
}

type Document struct {
	Id                      bson.ObjectId              `bson:"_id,omitempty"`
	DocId                   string                     `bson:"doc_id"`
	AppId                   string                     `bson:"app_id"`
	Title                   string                     `bson:"title"`
	Pages                   []*Page                    `bson:"pages"`
	Bizdata                 map[string]json.RawMessage `bson:"bizdata"`
	Operator                Operator                   `bson:"operator"`
	Status                  string                     `bson:"status,omitempty"`
	DecompressionFailedKeys []string                   `bson:"decompression_failed_keys,omitempty"`
}

type Page struct {
	Key     string                     `bson:"key"`
	Bizdata map[string]json.RawMessage `bson:"bizdata"`
	PageNo  int32                      `bson:"page_no"`
	Status  string                     `bson:"status,omitempty"`
}

type Operator struct {
	CreatedBy    string    `bson:"created_by"`
	CreatedAt    time.Time `bson:"created_at"`
	CreatedByApp string    `bson:"created_by_app"`
	UpdatedBy    string    `bson:"updated_by"`
	UpdatedAt    time.Time `bson:"updated_at"`
	UpdatedByApp string    `bson:"updated_by_app"`
}

func main() {
	var (
		results []*Document
	)

	session, err := mgo.Dial(DBEndPoint)
	if err != nil {
		log.Printf("Failed to dail db %s", err)
		return
	}

	c := session.DB(DBName).C(collectionName)
	err = c.Find(nil).All(&results)
	for _, result := range results {
		showDocument(result)
	}
}

func showDocument(doc *Document) {
	log.Println("DocId:", doc.DocId)
	log.Println("AppId:", doc.AppId)
	log.Println("Title:", doc.Title)
	log.Println("Bizdata:", doc.Bizdata)
	log.Println("status", doc.Status)

	for _, page := range doc.Pages {
		log.Println("page.key:", page.Key)
		log.Println("page.pageNo", page.PageNo)
		// log.Println("bizdata", string(page.Bizdata))

		for k, v := range page.Bizdata {
			// 	log.Println("pageNo:", page.PageNo)
			// 	var rawtest string
			// 	json.Unmarshal(v, &rawtest)
			// 	log.Println("k:", k, "RawJson:", rawtest)
			// fmt.Println("pageNo:", page.PageNo, "tt:", tt)

			switch k {
			case "aa":
				log.Println("k:", k, "RawJson is int:", v)
			case "bb":
				log.Println("k:", k, "RawJson is bool:", v)
			case "cc":
				log.Println("k:", k, "RawJson is string:", string(v))
			}
		}
	}
}
