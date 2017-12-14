package main

import "fmt"

// func main() {

// 	var jsonBlob = `[
//     "aaa",
//     "bbb"
// ]`)
// 	var animals []string
// 	err := json.Unmarshal(jsonBlob, &animals)
// 	if err != nil {
// 		fmt.Println("error:", err)
// 	}
// 	fmt.Printf("%+v", animals)
// }

// type test struct {
// 	justtest []string
// }

// func main() {
// 	var mytest = test{}
// 	mytest.justtest = append(mytest.justtest, "aa")
// 	mytest.justtest = append(mytest.justtest, "aa")
// 	mytest.justtest = append(mytest.justtest, "aa")
// 	fmt.Println(len(mytest.justtest))
// 	fmt.Println(mytest.justtest[0])
// }

// func main() {
// 	type ColorGroup struct {
// 		ID     int      `json: "id_num"`
// 		Name   string   `json: "name"`
// 		Colors []string `json: "colors"`
// 	}

// 	group := ColorGroup{
// 		ID:     1,
// 		Name:   "Reds",
// 		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
// 	}

// 	b, err := json.Marshal(group)
// 	if err != nil {
// 		fmt.Println("error:", err)
// 	}
// 	os.Stdout.Write(b)
// }

// func main() {

// 	// for {
// 	// 	in := bufio.NewReader(os.Stdin)
// 	// 	fmt.Println(in.Buffered())
// 	// 	in.WriteTo(os.Stdout)
// 	// 	// fmt.Println("11")
// 	// }
// 	// mybuf := bytes.NewBuffer([]byte("沈亚军"))
// 	// for {
// 	// 	r, size, e := mybuf.ReadRune()
// 	// 	fmt.Println("r ", string(r), "size ", size, "e ", e)
// 	// 	fmt.Println(mybuf.Len())
// 	// 	if e == io.EOF {
// 	// 		break
// 	// 	}
// 	// }
// 	a := NewBitset()
// 	a.set(1)
// 	fmt.Println("a is", a)
// }

// type bitset uint32

// func NewBitset() bitset {
// 	var b bitset
// 	b = 0
// 	return b
// }

// func (b bitset) set(offset int) error {
// 	if offset > 32 {
// 		return fmt.Errorf("The array bounds")
// 	}

// 	var a uint32
// 	a = 1
// 	tmp := a << offset
// 	b = b | tmp
// 	return nil
// }

//
type maptest struct {
	a map[string]string
	b int
}

const mytest = 8

func main() {
	// var config atomic.Value
	// config.Store([]string{"a", "b"})
	// value := config.Load()
	// if value == nil {
	// 	fmt.Println("hah")
	// } else {
	// 	fmt.Println(value.([]string))
	// 	config.Store(nil)
	// }
	const docId = "UDMP-cc0287cff88261ee6180790104860e1f440455a4ee-97dcfa02-20170307044708-00000001"
	const num = "13323"
	fmt.Println("doc len is ", len(docId))
	fmt.Println("num len is", num[1])
	var name = maptest{}
	if name.a == nil {
		fmt.Println("haha")
	}
	name.a = make(map[string]string)
	if name.a == nil {
		fmt.Println("haha1")
	}
	name.a["a"] = "b"
	fmt.Println("...", name)
	mytest := []string{"a", "b", "c"}
	fmt.Println(mytest[1:])

	var mytest1 = make(map[string]int)
	mytest1["a"] = 3
	mytest1["a"] = 4
	fmt.Println("mytest1", mytest1)
	// test(2)
}

func test(a int) {
	if a == 0 {
		a = mytest
	}
	fmt.Println(a)
}
