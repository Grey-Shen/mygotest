package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type ESBQuery struct {
	XMLName xml.Name      `xml:"service"`
	Version string        `xml:"version,attr"`
	SysHead SYSHEADStruct `xml:"SYS_HEAD"`
	AppHead AppHeadStruct `xml:"APP_HEAD"`
	Body    BodyStruct    `xml:"BODY"`
}

type SYSHEADStruct struct {
	ServiceCode   ServiceCodeStruct
	ServiceScene  ServiceSceneStruct
	ConsumerID    ConsumerIDStruct
	TranDate      TranDateStruct
	TranTimeStamp TranTimeStampStruct
}

type AppHeadStruct struct {
	BussSeqNo BussSeqNoStruct
}

type BodyStruct struct {
	ObjectID ObjectIDStruct
	DocID    DocIDstruct
}

type ValueStruct struct {
	Attr  string `xml:"attr,attr"`
	Value string `xml:",chardata"`
}

type ServiceCodeStruct struct {
	XMLName xml.Name `xml:"SERVICE_CODE"`
	ValueStruct
}

type ServiceSceneStruct struct {
	XMLName xml.Name `xml:"SERVICE_SCENE"`
	ValueStruct
}

type ConsumerIDStruct struct {
	XMLName xml.Name `xml:"CONSUMER_ID"`
	ValueStruct
}

type TranDateStruct struct {
	XMLName xml.Name `xml:"TRAN_DATE"`
	ValueStruct
}

type TranTimeStampStruct struct {
	XMLName xml.Name `xml:"TRAN_TIMESTSMP"`
	ValueStruct
}

type BussSeqNoStruct struct {
	XMLName xml.Name `xml:"BULL_SEQ_NO"`
	ValueStruct
}

type ObjectIDStruct struct {
	XMLName xml.Name `xml:"OBJECT_ID"`
	ValueStruct
}

type DocIDstruct struct {
	XMLName xml.Name `xml:"DOC_ID"`
	ValueStruct
}

//////////////////////////

type ESBResponse struct {
	XMLName xml.Name          `xml:"service"`
	Version string            `xml:"version,attr"`
	SysHeaD RespSysHeadStruct `xml:"SYS_HEAD"`
	Body    RespBodyStruct    `xml:"BODY"`
}

type RespSysHeadStruct struct {
	ConsumerID ConsumerIDStruct
}

type RespBodyStruct struct {
	RespCode RespCodeStruct
	DealMesg DealMsgStruct
	FilePath FilePathStruct
	FileName FileName1Struct
}

type RespCodeStruct struct {
	XMLName xml.Name `xml:"RSP_CODE"`
	ValueStruct
}

type DealMsgStruct struct {
	XMLName xml.Name `xml:"DEAL_MESG"`
	ValueStruct
}

type FilePathStruct struct {
	XMLName xml.Name `xml:"FILE_PATH"`
	ValueStruct
}

type FileName1Struct struct {
	XMLName xml.Name `xml:"FILE_NAME1"`
	ValueStruct
}

func main() {
	service := ESBQuery{}
	service.Version = "2.0"
	service.SysHead.ServiceCode.Attr = "s,15"
	service.SysHead.ServiceCode.Value = "11003000026"

	service.SysHead.ServiceScene.Attr = "s,2"
	service.SysHead.ServiceScene.Value = "02"

	service.SysHead.ConsumerID.Attr = "s,6"
	service.SysHead.ConsumerID.Value = "011701"

	service.SysHead.TranDate.Attr = "s,8"
	service.SysHead.TranDate.Value = "20180927"

	service.SysHead.TranTimeStamp.Attr = "s,6"
	service.SysHead.TranTimeStamp.Value = "103912"

	service.AppHead.BussSeqNo.Attr = "s,26"
	service.AppHead.BussSeqNo.Value = "fake"

	service.Body.ObjectID.Attr = "s.6"
	service.Body.ObjectID.Value = "05091"

	service.Body.DocID.Attr = "s,20"
	service.Body.DocID.Value = "fake1"

	if m, err2 := xml.MarshalIndent(service, "", "\t"); err2 != nil {
		panic("xml.MarshalIndent FAILED: " + err2.Error())
	} else {
		xmlheader := xml.Header
		m = append([]byte(xmlheader), m...)
		fmt.Printf("%s", m)
	}
	////// unmarshal
	fmt.Println("======================= unmarshal")

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
	v := ESBResponse{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println(v)
}
