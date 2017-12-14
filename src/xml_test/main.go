package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

// type CentAuthroty struct {
// 	XMLName      xml.Name     `xml:"CentAuthroty"`        // 指定最外层的标签为config
// 	FilesCount   int          `xml:"TransportFilesCount"` // 读取smtpServer配置项，并将结果保存到SmtpServer变量中
// 	Status       int          `xml:"status"`
// 	SenderPasswd string       `xml:"senderPasswd"`
// 	TransferList TransferList `xml:"TransferList"` // 读取receivers标签下的内容，以结构方式获取
// }

type CentAuthroty struct {
	XMLName    xml.Name   `xml:"CentAuthroty"` // 指定最外层的标签为config
	FilesCount FileCountS // 读取smtpServer配置项，并将结果保存到SmtpServer变量中
	Status     int        `xml:"status"`
}

type FileCountS struct {
	XMLName xml.Name `xml:"TransportFilesCount"`
	ValueStruct
}

// type ValueStruct struct {
// 	Attr  string `xml:"attr,attr"`
// 	Value string `xml:",chardata"`
// }

type TransferList struct {
	Transfers []Transfer `xml:"Transfer"`
}

type Transfer struct {
	TxLogNo       string    `xml:"TX_LOG_NO"`
	TxDate        string    `xml:"TX_DATE1"`
	CustID        string    `xml:"CUST_ID"`
	MachineBranch string    `xml:"MACHINE_BRANCH"`
	MachineTeller string    `xml:"MACHINE_TELLER1"`
	TxType        int       `xml:"TX_TYPE"`
	FileNames     FileNames `xml:"FileNames"`
}

type FileNames struct {
	FileName []string `xml:"FileName"`
}

func main11() {
	file, err := os.Open("test1.xml") // For read access.
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	v := CentAuthroty{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println(v)
}
