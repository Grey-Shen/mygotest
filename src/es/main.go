package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"metric/funcs"
	"metric/model"
	"net/http"
	"strings"
	"time"

	"github.com/spf13/viper"
)

func LogCollect() {
	for {
		logCollect()
		time.Sleep(time.Duration(5) * time.Second)
	}
}

func logCollect() {

	url := "http://" + viper.GetString("es.url") + "/_search"
	reqBody := new(model.RequestValue)
	lte := time.Now().UTC().Format("2006-01-02T15:04:05-0700")
	fm, _ := time.ParseDuration("-1m")
	gte := time.Now().UTC().Add(fm).Format("2006-01-02T15:04:05-0700")
	logtime := model.Logtime{Gte: gte, Lte: lte}
	//logtime := model.Logtime{Gte:"2017-10-20T09:45:35.015+0800", Lte:"2017-10-24T09:45:35.015+0800"}
	filter := model.Filter{
		model.Siglefilter{"term": model.Siglefilter{"loglvl.keyword": "ERROR"}},
		model.Siglefilter{"range": model.Siglefilter{"logtime": logtime}},
	}

	reqBody.Source.Excludes = []string{"offset", "*type", "beat", "*timestamp"}
	reqBody.Query.Bool.Filter = filter
	reqJson, err := json.Marshal(reqBody)
	if err != nil {
		panic(err)
	}

	payload := strings.NewReader(string(reqJson))

	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		panic(err)
	}

	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	log.Println(string(body))

	response, err := simplejson.NewJson(body)
	//response, err := simplejson.NewJson([]byte(`{"took":19,"timed_out":false,"_shards":{"total":5,"successful":5,"failed":0},"hits":{"total":2,"max_score":0,"hits":[{"_index":"test_log","_type":"doc","_id":"AV_DibsASlqUka4Ai8T1","_score":0,"_source":{"method":"-","loglvl":"ERROR","module":"piaowu","source":"/var/log/oto_saas/piaowu.log","message":"2017-11-16T14:34:52.964+0800 ERROR piaowu - [Catch Error on POST /piaowu/saveOrder, body: {\"customerUserId\":\"test_long7\",\"userPhone\":\"13512119093\",\"quantity\":1,\"dispatchWay\":3,\"ticketId\":\"2033021\",\"channel\":\"xishiqu\",\"contact\":{\"name\":\"我才\",\"phone\":\"13512119099\"},\"contactId\":\"\",\"activityCode\":\"20170823006\",\"eventId\":230183,\"couponId\":\"\",\"activityId\":\"\",\"appCode\":\"me\",\"customerId\":\"209\"}]\n2017-11-16T14:34:52.965+0800 ERROR piaowu - [Catch Error: Error: 没有找到该票或该票已售罄\n    at Promise.all.then.then.result (/usr/src/app/oto_saas_piaowu/source/lib/xishiqu/order.js:104:23)\n    at process._tickDomainCallback (internal/process/next_tick.js:135:7)]\n2017-11-16T14:34:52.965+0800 ERROR piaowu - [Catch Error stack: Error: 没有找到该票或该票已售罄\n    at Promise.all.then.then.result (/usr/src/app/oto_saas_piaowu/source/lib/xishiqu/order.js:104:23)\n    at process._tickDomainCallback (internal/process/next_tick.js:135:7)]\n2017-11-16T14:34:52.966+0800 INFO piaowu - [Error: 没有找到该票或该票已售罄\n    at Promise.all.then.then.result (/usr/src/app/oto_saas_piaowu/source/lib/xishiqu/order.js:104:23)\n    at process._tickDomainCallback (internal/process/next_tick.js:135:7)]","logtime":"2017-11-16T14:34:52.964+0800"}},{"_index":"test_log","_type":"doc","_id":"AV_DiYCSMh8gBf2s33uS","_score":0,"_source":{"method":"-","loglvl":"ERROR","module":"piaowu","source":"/var/log/oto_saas/piaowu.log","message":"2017-11-16T14:34:33.244+0800 ERROR piaowu - [Catch Error on POST /piaowu/saveOrder, body: {\"customerUserId\":\"test_long7\",\"userPhone\":\"13512119093\",\"quantity\":1,\"dispatchWay\":3,\"ticketId\":\"2039544\",\"channel\":\"xishiqu\",\"contact\":{\"name\":\"我才\",\"phone\":\"13512119099\"},\"contactId\":\"\",\"activityCode\":\"20170224015\",\"eventId\":190616,\"couponId\":\"\",\"activityId\":\"\",\"appCode\":\"me\",\"customerId\":\"209\"}]\n2017-11-16T14:34:33.245+0800 ERROR piaowu - [Catch Error: Error: 没有找到该票或该票已售罄\n    at Promise.all.then.then.result (/usr/src/app/oto_saas_piaowu/source/lib/xishiqu/order.js:104:23)\n    at process._tickDomainCallback (internal/process/next_tick.js:135:7)]\n2017-11-16T14:34:33.245+0800 ERROR piaowu - [Catch Error stack: Error: 没有找到该票或该票已售罄\n    at Promise.all.then.then.result (/usr/src/app/oto_saas_piaowu/source/lib/xishiqu/order.js:104:23)\n    at process._tickDomainCallback (internal/process/next_tick.js:135:7)]\n2017-11-16T14:34:33.246+0800 INFO piaowu - [Error: 没有找到该票或该票已售罄\n    at Promise.all.then.then.result (/usr/src/app/oto_saas_piaowu/source/lib/xishiqu/order.js:104:23)\n    at process._tickDomainCallback (internal/process/next_tick.js:135:7)]","logtime":"2017-11-16T14:34:33.244+0800"}}]}}`))

	if err != nil {
		panic(err)
	}

	total, _ := response.Get("hits").Get("total").Int()
	if total != 0 {
		hits := response.Get("hits").Get("hits")
		logMsg, err := json.MarshalIndent(hits, "", "        ")
		if err != nil {
			panic(err)
		}
		funcs.SendMail([]string{"wuyachao@boluome.com"}, string(logMsg))
		fmt.Println(string(logMsg))
	}
}
