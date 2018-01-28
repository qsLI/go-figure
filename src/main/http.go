package main

import (
	"net/http"
	"log"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

var f interface{}

type Speaker struct {
	FirstName   string
	LastName    string
	Category    string
	Title       string
	Image       string
	Link        string
	Bio         string
	Description string
}

type Speakers struct {
	Speakers []Speaker `json: "speakers"`
}


func main() {
	uri := "https://gist.githubusercontent.com/qsLI/cbc4311d186f9e9c13d01261c5318d95/raw/af7bcff168eeaeb6eccc4f3f67bdc5410847e29f/test.json"

	resp, err := http.Get(uri)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	//// 输出header信息
	//fmt.Println(resp.Status)
	//for k, v := range resp.Header {
	//	fmt.Println(k, " = ", v)
	//}

	bytes, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		fmt.Println("read response body error", e)
		return
	}

	var speakers Speakers
	json.Unmarshal(bytes, &speakers)

	fmt.Println(speakers)

	testMarshal()


	//json.Unmarshal(bytes, &f)
	//m := f.(map[string]interface{})
	//
	//for k, v := range m {
	//	switch vv := v.(type) {
	//	case string:
	//		fmt.Println(k, "is string", vv)
	//	case float64:
	//		fmt.Println(k, "is float64", vv)
	//	case []interface{}:
	//		fmt.Println(k, "is an array:")
	//		for i, u := range vv {
	//			fmt.Println(i, u)
	//		}
	//	default:
	//		fmt.Println(k, "is of a type I don't know how to handle")
	//	}
	//}

}

func testMarshal() {
	var speakers Speakers
	speaker := Speaker{
		"kevein",
		"Leo",
		"dev",
		"rd",
		"http://test.jsp",
		"http://test.jsp",
		"",
		"",
	}
	speakers.Speakers = []Speaker{speaker}
	result, _ := json.Marshal(speakers)
	fmt.Println(string(result))
}
