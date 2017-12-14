package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
)

/***************************************************************************************************
    Global Variables
***************************************************************************************************/
var gCurCookies []*http.Cookie
var gCurCookieJar *cookiejar.Jar

/***************************************************************************************************
    Functions
***************************************************************************************************/
//do init before all others
func initAll() {
	gCurCookies = nil
	//var err error;
	gCurCookieJar, _ = cookiejar.New(nil)

}

//get url response html
func getUrlRespHtml(url string) string {
	fmt.Printf("getUrlRespHtml, url=%s", url)

	var respHtml string = ""

	httpClient := &http.Client{
		CheckRedirect: nil,
		Jar:           gCurCookieJar,
	}

	// httpResp, err := httpClient.Get("http://example.com")
	// // ...

	httpReq, err := http.NewRequest("GET", url, nil)
	// ...
	//httpReq.Header.Add("If-None-Match", `W/"wyzzy"`)
	httpResp, err := httpClient.Do(httpReq)
	// ...

	//httpResp, err := http.Get(url)
	//gLogger.Info("http.Get done")
	if err != nil {
		fmt.Printf("http get url=%s response error=%s\n", url, err.Error())
	}
	fmt.Printf("httpResp.Header=%s", httpResp.Header)
	fmt.Printf("httpResp.Status=%s", httpResp.Status)

	defer httpResp.Body.Close()
	// gLogger.Info("defer httpResp.Body.Close done")

	body, errReadAll := ioutil.ReadAll(httpResp.Body)
	//gLogger.Info("ioutil.ReadAll done")
	if errReadAll != nil {
		fmt.Printf("get response for url=%s got error=%s\n", url, errReadAll.Error())
	}
	//gLogger.Debug("body=%s\n", body)

	//gCurCookies = httpResp.Cookies()
	//gCurCookieJar = httpClient.Jar;
	gCurCookies = gCurCookieJar.Cookies(httpReq.URL)
	//gLogger.Info("httpResp.Cookies done")

	//respHtml = "just for test log ok or not"
	respHtml = string(body)
	//gLogger.Info("httpResp body []byte to string done")

	return respHtml
}

func dbgPrintCurCookies() {
	var cookieNum int = len(gCurCookies)
	fmt.Printf("cookieNum=%d", cookieNum)
	for i := 0; i < cookieNum; i++ {
		var curCk *http.Cookie = gCurCookies[i]
		//gLogger.Info("curCk.Raw=%s", curCk.Raw)
		fmt.Printf("------ Cookie [%d]------", i)
		fmt.Printf("Name\t=%s", curCk.Name)
		fmt.Printf("Value\t=%s", curCk.Value)
		fmt.Printf("Path\t=%s", curCk.Path)
		fmt.Printf("Domain\t=%s", curCk.Domain)
		fmt.Printf("Expires\t=%s", curCk.Expires)
		fmt.Printf("RawExpires=%s", curCk.RawExpires)
		fmt.Printf("MaxAge\t=%d", curCk.MaxAge)
		fmt.Printf("Secure\t=%t", curCk.Secure)
		fmt.Printf("HttpOnly=%t", curCk.HttpOnly)
		fmt.Printf("Raw\t=%s", curCk.Raw)
		fmt.Printf("Unparsed=%s", curCk.Unparsed)
	}
}

func main() {
	initAll()

	//step1: access baidu url to get cookie BAIDUID
	fmt.Println("====== 步骤1：获得BAIDUID的Cookie ======")
	var baiduMainUrl string = "http://www.baidu.com/"
	fmt.Printf("baiduMainUrl=%s", baiduMainUrl)
	respHtml := getUrlRespHtml(baiduMainUrl)
	fmt.Printf("respHtml=%s", respHtml)
	dbgPrintCurCookies()

	//check cookie
	//...

	//step2: login, pass paras, extract resp cookie
	fmt.Println("====== 步骤2：提取login_token ======")
	// if bBaiduidCookieExist {
	//https://passport.baidu.com/v2/api/?getapi&class=login&tpl=mn&tangram=true
	var getapiUrl string = "https://passport.baidu.com/v2/api/?getapi&class=login&tpl=mn&tangram=true"
	var getApiRespHtml string = getUrlRespHtml(getapiUrl)
	fmt.Printf("getApiRespHtml=%s", getApiRespHtml)
	dbgPrintCurCookies()
	// }
}
