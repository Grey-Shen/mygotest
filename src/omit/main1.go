package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//type KongResult struct {
//  Total int `json:"total"`
//  Next  int `json:"next"`
//  Data []struct {
//      StripURI     bool     `json:"strip_uri"`
//      Name         string   `json:"name"`
//      UpstreamURL  string   `json:"upstream_url"`
//      Uris         []string `json:"uris"`
//      Hosts        []string `json:"hosts"`
//      PreserveHost bool     `json:"preserve_host"`
//  } `json:"data"`
//}

type SingleService struct {
	Name         string   `json:"name"`
	Uris         []string `json:"uris"`
	Hosts        []string `json:"hosts,omitempty"`
	UpstreamURL  string   `json:"upstream_url"`
	PreserveHost bool     `json:"preserve_host"`
	StripURI     bool     `json:"strip_uri"`
}

func main() {

	env := flag.String("env", "dev", "dev,stg,pro")

	flag.Parse()

	var service SingleService
	var url string
	result := make(map[string]SingleService)

	switch *env {
	case "dev":
		url = "http://192.168.1.12:8001/apis/"
	default:
		log.Fatal("env error:", *env)
	}

	//h5 response
	h5_resp, err := http.Get(url + "h5-api-waimai")
	if err != nil {
		log.Println("connect to gateway server failed:", err)
	}

	defer h5_resp.Body.Close()
	h5_body, err := ioutil.ReadAll(h5_resp.Body)

	//h5_body{}
	json.Unmarshal(h5_body, &service)

	result["h5"] = service
	log.Println(service.Uris)
	log.Println(service)
	//api response
	api_resp, err := http.Get(url + "api-waimai")
	if err != nil {
		log.Println("connect to gateway server failed:", err)
	}
	defer api_resp.Body.Close()
	api_body, err := ioutil.ReadAll(api_resp.Body)

	json.Unmarshal(api_body, &service)
	log.Println(service.Uris)
	log.Println(service)
	result["api"] = service

	api_res, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println("api json decode error:", err)
	}
	fmt.Println(string(api_res))
}
