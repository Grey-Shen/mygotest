package main

import (
	"fmt"
	"log"

	"github.com/globalsign/mgo"
)

type User struct {
	Id    string `bson:"_id"`
	Email string `bson:"email"`
}

func main() {
	// Database
	dbs, err := mgo.Dial("mongodb://localhost/")
	if err != nil {
		panic(err)
	}
	// Collections
	uc := dbs.Clone().DB("").C("users")
	defer dbs.Clone().DB("").Session.Close()
	users := make([]interface{}, 2)
	user := User{
		Id:    "999",
		Email: "111@exmple.com",
	}

	users = append(users, user)

	user1 := User{
		Id:    "888",
		Email: "111@exmple.com",
	}

	users = append(users, user1)

	user2 := User{
		Id:    "100",
		Email: "111@exmple.com",
	}

	users = append(users, user2)

	bulk := uc.Bulk()
	bulk.Unordered()
	bulk.Insert(users...)
	res, bulkErr := bulk.Run()
	if bulkErr != nil {
		if bulkError, ok := bulkErr.(*mgo.BulkError); !ok {
			// not a bulk error
		} else {
			fmt.Println("========= ", bulkError.Cases(), len(bulkError.Cases()))
			for _, c := range bulkError.Cases() {
				fmt.Println("=======", c.Err)
				if mgo.IsDup(c.Err) {
					fmt.Println("==duplicate==", c.Err)
					fmt.Println("===n==", c.Index)
				}
			}
		}
	}

	log.Println("==== res: ", res)
}
