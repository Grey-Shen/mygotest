package main

import (
	"encoding/json"
	"fmt"
	"log"
)

var myraw = json.RawMessage(`{
    "title": "...",
    "bizdata": {
        "type": true,
        "city": 88.8
    },
    "created_by": "柯南",
    "pages": [
        {
            "bizdata": {
                "category": "侦探",
                "user": "毛利小五郎"
            }
        },
        {
            "bizdata": {
                "user": "阿笠博士",
                "年龄": "50"
            }
        },
        {
        }
    ]
}`)

type DocumentCreateParams struct {
	Title     string                     `json:"title"`
	Bizdata   map[string]json.RawMessage `json:"bizdata"`
	CreatedBy string                     `json:"created_by" binding:"required"`
	Pages     []*PageCreateParams        `json:"pages" binding:"required,min=1,dive"`
}

type PageCreateParams struct {
	Bizdata map[string]json.RawMessage `json:"bizdata"`
}

func main() {
	var params *DocumentCreateParams
	json.Unmarshal(myraw, &params)
	showParams(params)
}

func showParams(params *DocumentCreateParams) {
	log.Println("title", params.Title)
	log.Println("createdby", params.CreatedBy)
	log.Println("bizdata:", params.Bizdata)
	for k, v := range params.Bizdata {
		switch k {
		case "type":
			var x bool
			json.Unmarshal(v, &x)
			fmt.Println("x:", x)
			// case "city":
			// 	var x bool
			// 	b_buf := bytes.NewBuffer(v)
			// 	binary.Read(b_buf, binary.LittleEndian, &x)
			// 	fmt.Println("k:", k, "city--:", x)
			// 	log.Println("k:", k, "city:", v)

			// var b bool = v.(bool)
			// fmt.Println("b:", b)
		}
	}

	for _, page := range params.Pages {
		fmt.Println("page_bizdata:", page.Bizdata)
	}
}
