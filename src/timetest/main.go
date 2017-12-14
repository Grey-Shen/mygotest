package main

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

const myformat = "2006-01-02"

var timetest time.Time

func main() {
	// d, err := parseDateString("2016-12-21 18:23:50")
	// if err != nil {
	// 	fmt.Println(err)
	// } e	lse {
	// 	fmt.Println("date is ", d)
	// }
	// testDate := time.Date(d.year, d.month, d.day, d.hour, d.minute, d.second, 0, time.Local)
	// fmt.Println("testDate", testDate)

	mytime, err := time.Parse(time.RFC3339, "2017-07-09T11:43:19+08:00")
	fmt.Println("nowtest", time.Now().Format(myformat))
	if mytime.IsZero() {
		fmt.Println("yes")
	}

	// mytime, err := time.Parse(time.RFC822Z, "02 Jan 06 15:04 -0700")
	// mytime, err := time.ParseInLocation(myformat, "2006-01-02 15:04:05", nil)
	if err != nil {
		fmt.Println("hah")
	}
	// year, month, day := mytime.Date()
	fmt.Println("mytime", mytime)
	// t := time.Date(2009, time.November, 10, 31, 8, 2, 0, time.UTC)
	// fmt.Printf("Go launched at %s\n", t.Local())

	fmt.Println("test", url.QueryEscape("2017-07-08T14:53:43 08:00"))
	var test = "hahaha"
	fmt.Println("test", test[0:3])

	fmt.Printf("%q\n", strings.SplitAfter("Too frequent, please wait for 5 seconds", " s"))

}

// type date struct {
// 	year   intfencode
// 	month  time.Month
// 	day    int
// 	hour   int
// 	minute int
// 	second int
// }

// func parseDateString(value string) (date, error) {
// 	var d date
// 	if n, err := fmt.Sscanf(value, "%d-%d-%d %d:%d:%d", &d.year, &d.month, &d.day, &d.hour, &d.minute, &d.second); err != nil {
// 		return d, fmt.Errorf("Invalid date format: %s:%s", value, err)
// 	} else if n != 6 {
// 		return d, fmt.Errorf("Invalid date format: %s", value)
// 	} else {
// 		return d, nil
// 	}
// }

// Run command unmarshaled: mgo.queryOp{collection:"tassadar_release.$cmd", query:mgo.pipeCmd{Aggregate:"app_api_call_stats", Pipeline:[]bson.M{bson.M{"$match":bson.M{"app_id":"2017-05-10 18:00:00"}}, bson.M{"$group":bson.M{"_id":bson.M{"appId":"$app_id", "apiName":"$api_name"}, "count":bson.M{"$sum":"$count"}}}}, Cursor:(*mgo.pipeCmdCursor)(0xc420280170), Explain:false, AllowDisk:false}, skip:0, limit:-1, selector:interface {}(nil), flags:0x0, replyFunc:(mgo.replyFunc)(0x14103d0), mode:2, options:mgo.queryWrapper{Query:interface {}(nil), OrderBy:interface {}(nil), Hint:interface {}(nil), Explain:false, Snapshot:false, ReadPreference:bson.D(nil), MaxScan:0, MaxTimeMS:0, Comment:""}, hasOptions:false, serverTags:[]bson.D(nil)}, result: bson.M{"waitedMS":0, "cursor":bson.M{"id":0, "ns":"tassadar_release.app_api_call_stats", "firstBatch":[]interface {}{}}, "ok":1}
