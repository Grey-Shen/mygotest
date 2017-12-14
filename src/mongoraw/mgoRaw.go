package main

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type GeneralRecord struct {
	Type string `bson:"type"`
}

type Person struct {
	Name  string `bson:"name"`
	Phone string `bson:"phone"`
}

type Company struct {
	Name string `bson:"company"`
	Boss string `bson:"boss"`
}

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	var rawRecords []*bson.Raw

	c := session.DB("test").C("people")
	if err := c.Find(nil).All(&rawRecords); err != nil {
		panic(err)
	} else {
		var general GeneralRecord
		var person Person
		var company Company

		for _, raw := range rawRecords {
			if err = raw.Unmarshal(&general); err != nil {
				panic(err)
			}
			fmt.Println("typetest", general.Type)
			switch general.Type {
			case "person":
				if err = raw.Unmarshal(&person); err != nil {
					panic(err)
				} else {
					fmt.Printf("Person: %#v\n", person)
				}
			case "company":
				if err = raw.Unmarshal(&company); err != nil {
					panic(err)
				} else {
					fmt.Printf("Company: %#v\n", company)
				}
			}
		}
	}
}
