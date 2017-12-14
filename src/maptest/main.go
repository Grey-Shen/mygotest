// package main

// import "fmt"

// func main() {
// 	testMap := make(map[int]string, 3)
// 	testMap[1] = "haha"
// 	if len(testMap) == 0 {
// 		fmt.Println("hah")
// 	} else if testMap[1] == "haha" {
// 		fmt.Println("yes")
// 	}
// }

package main

import (
	"fmt"
	"log"
	"reflect"
)

func main() {
	// a := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e", 6: "g"}
	// for k, v := range a {
	// 	fmt.Printf("k: %d, v: %s\n", k, v)
	// }
	var b interface{}
	b = 1
	fmt.Errorf("invalid b: %#v", b)
	t := fmt.Sprintf("invalid b %#v", b)
	fmt.Println("test", t)
	s, ok := b.(string)
	if !ok {
		log.Printf("got data of type %T but wanted string", b)
	}
	log.Println("test", s)

	// decoded, err := base64.URLEncoding.DecodeString("eyJzY29wZSI6InVkbXAtc3RvcmFnZTpVRE1QLTU4MTFkMjhjOGRjOTA3OGJlOWM1Yjk1NmI2OWU4Y2MwYTc3NGVlMjg4YS0zYTc1YzM1ZS0yMDE3MTExNjAxMjYyMi0wMDAwMDAwMi8yIiwiZGVhZGxpbmUiOjE1MTA4ODE5ODIsImluc2VydE9ubHkiOjEsImRldGVjdE1pbWUiOjEsInNhdmVLZXkiOiJVRE1QLTU4MTFkMjhjOGRjOTA3OGJlOWM1Yjk1NmI2OWU4Y2MwYTc3NGVlMjg4YS0zYTc1YzM1ZS0yMDE3MTExNjAxMjYyMi0wMDAwMDAwMi8yIiwiY2FsbGJhY2tVcmwiOiJodHRwOi8vMTAuMTQuNTIuODI6NDAwMC9kb2N1bWVudHMvY2FsbGJhY2svY3JlYXRlL1VETVAtNTgxMWQyOGM4ZGM5MDc4YmU5YzViOTU2YjY5ZThjYzBhNzc0ZWUyODhhLTNhNzVjMzVlLTIwMTcxMTE2MDEyNjIyLTAwMDAwMDAyIiwiY2FsbGJhY2tCb2R5Ijoia2V5PSQoa2V5KSIsImNhbGxiYWNrQm9keVR5cGUiOiJhcHBsaWNhdGlvbi94LXd3dy1mb3JtLXVybGVuY29kZWQifQ==")
	// if err != nil {
	// 	fmt.Println("decode error:", err)
	// 	return
	// }
	// fmt.Println(string(decoded))

	list("aa", "bb")
}

func list(tt ...string) {
	for _, v := range tt {
		fmt.Println("tttest", v)
	}
}

func Contains(s interface{}, elem interface{}) bool {
	arr := reflect.ValueOf(s)

	if arr.Kind() == reflect.Slice || arr.Kind() == reflect.Array {
		for i := 0; i < arr.Len(); i++ {
			if arr.Index(i).Interface() == elem {
				return true
			}
		}
	}

	return false
}

// db.async_tasks.find({"$and": [{"$or": [{"status": "executing"},{"status": "scheduled"}]},"task_type": "pfop"]})
