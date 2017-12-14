// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	// u, err := url.Parse("http://bing.com/search?q=dotnet")
// 	// u, err := url.Parse("//bing.com/searfdfdch?q=dotnet")
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// // u.Scheme = "https"
// 	// // u.Host = "google.com"
// 	// // q := u.Query()
// 	// // q.Set("q", "golang")
// 	// // u.RawQuery = q.Encode()
// 	// fmt.Println(u.String())

// 	test := "-hello- 世界"
// 	fmt.Println(strconv.QuoteToASCII(test))
// 	// appId := "fdfdf"
// 	fmt.Println(fmt.Sprintf(`-%s-`, test))
// }

package main

import (
	"fmt"
	"regexp"
)

const docIdExpr = `UDMP-[0-9a-f]{42}-[0-9a-f]{8}-\d{14}-[0-9a-f]{8}`
const packageIdExpr = `UDMP-PACKAGES-[0-9a-f]{42}-[0-9a-f]{8}-\d{14}`
const reserveExpr = `reserve_upload/.{8}`

func main() {
	// myregexp := regexp.MustCompile(`^UDMP-[0-9a-f]{42}-[0-9a-f]{8}-\d{14}-[0-9a-f]{8}$`)
	// if err != nil {
	// 	fmt.Println("Failed to get regexp err", err)
	// 	return
	// }

	docIdRegexp := regexp.MustCompile(docIdExpr)

	packageIdRegexp := regexp.MustCompile(packageIdExpr)
	var documentquery = "/documents/UDMP-999d11f757e594603681d73f8b26d15515953b0f09-97dcfa02-20170307044706-00000001/2"
	pageNoRegexp := regexp.MustCompile(`/\d`)
	tmp := docIdRegexp.ReplaceAllString(documentquery, ":docId")
	fmt.Println("tmp", tmp)
	tmp = pageNoRegexp.ReplaceAllString(tmp, "/:pageNo")
	fmt.Println("tmp", tmp)
	tmp = packageIdRegexp.ReplaceAllString(tmp, "/:package_id")
	fmt.Println("tmp", string(tmp))

	reserveRegexp := regexp.MustCompile(reserveExpr)

	resrveTest := "/documents/reserve_upload/97dcfa02"
	tmp = reserveRegexp.ReplaceAllString(resrveTest, "reserve_upload/:appId")
	fmt.Println("reservetest", tmp)
	// q := strings.Split(documentquery, "/")
	// var result string
	// for _, v := range q {
	// 	if v != "" {
	// 		fmt.Println("v is", v)
	// 		if myregexp.MatchString(v) {
	// 			fmt.Println("match ok")
	// 			result = result + "/" + ":docId"
	// 		} else {
	// 			result = result + "/" + v
	// 		}
	// 	}

	// }
	// fmt.Println("result", result)

	// re := regexp.MustCompile("a(x*)b")
	// fmt.Println(re.ReplaceAllString("-ab-axxb-", "T"))
	// fmt.Println(re.ReplaceAllString("-ab-axxb-", "$1"))
	// fmt.Println(re.ReplaceAllString("-ab-axxb-", "$1W"))
	// fmt.Println(re.ReplaceAllString("-ab-axxb-", "${1}W"))

	urltest := "https://bing.com/search?q=dotnet"
	re := regexp.MustCompile(`^(http|https)://`)
	tt := re.ReplaceAllString(urltest, "")
	fmt.Println(tt)

}
