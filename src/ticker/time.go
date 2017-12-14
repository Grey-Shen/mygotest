package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// func main() {

// 	// s := "1479288956403"
// 	// 1492910846668781866
// 	// 14836876538273930
// 	// t, err := time.ParseDuration(s)
// 	// if err != nil {
// 	// 	fmt.Println("err ", err)
// 	// }
// 	// fmt.Println(t.Seconds())
// 	// fmt.Println(t.Nanoseconds())
// 	// fmt.Println(t.Hours())
// 	// 1492911443
// 	second := 14836876538273930 / 10000000
// 	fmt.Println("second", second)
// 	t := time.Unix(int64(second), 0)
// 	fmt.Println(t.Format("2006-01-02 15:04:05"))
// 	// stringtest := []string{"aa", "bb", "cc"}
// 	// fmt.Fprintf(os.Stdout, stringtest)
// 	// s.Format("2006-01-02 15:04:05")
// 	// fmt.Println("time: ", time.Now().UnixNano())
// 	curTime := time.Date(2107, time.May, 10, 12, 10, 10, 0, time.UTC)
// 	fmt.Println("timead", curTime.Add(1*time.Minute))
// 	fmt.Println("timeunix", time.Now().Unix()+7200)
// 	a := []byte("hell")
// 	fmt.Println("aa", string(a))

// 	// fmt.Printf("[%q]", strings.Trim(" !!! Achtung! Achtung! !!! ", "! "))
// 	aa := "http://fakedomain/fakepath1?e=1495620558&token=bef125afc8e21650d92bb141749bce7aeec538fafaf274f88e:XBc_48idjQqNg_YkHcbC-uF7sB4&life"
// 	myurl, _ := url.Parse(aa)
// 	aint, _ := strconv.Atoi(myurl.Query().Get("e"))
// 	myurl.Scheme = "https"
// 	aa = myurl.String()
// 	fmt.Println("https_aa", aa)

// 	fmt.Println("11", aint)
// 	fmt.Println("22", myurl.Query().Get("life"))

// 	type mytest struct {
// 		a int
// 		b int
// 	}

// 	hahatest := mytest{
// 		a: 5,
// 	}
// 	fmt.Println("hahatest", hahatest)
// 	// re := regexp.MustCompile("a(x*)b(y|z)c")
// 	// fmt.Printf("%q\n", re.FindStringSubmatch("-axxxbyc-"))
// 	// fmt.Printf("%q\n", re.FindStringSubmatch("-abzc-"))
// 	// re = regexp.MustCompile("fo.?")
// 	// fmt.Printf("%q\n", re.FindString("seafood"))
// // 	// fmt.Printf("%q\n", re.FindString("meat"))
// }

func main() {
	const url = "http://78re52.com1.z0.glb.clouddn.com/resource/Ship.jpg?imageView2/2/w/200/h/200|saveas/cWluaXUtZGV2ZWxvcGVyOlNoaXAtdGh1bWItMjAwLmpwZw==/sign/bcgojLbLKTsTlhm3XFMYq0cn3lW2G3NAuJYXZDDf:jGo09Pmq5vyG4c-rRb4qF3_dH1="
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("bodytest", string(bytes))
	}

}
