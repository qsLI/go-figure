package main

import (
	"net/http"
	"log"
	"fmt"
	"io/ioutil"
	"encoding/json"
)


const (
	url = "https://gist.githubusercontent.com/qsLI/cbc4311d186f9e9c13d01261c5318d95/raw/4719a2d1c93fcb166da2bdace546dabd8135b4eb/test.json"
)
// curl -s "https://gist.githubusercontent.com/qsLI/cbc4311d186f9e9c13d01261c5318d95/raw/4719a2d1c93fcb166da2bdace546dabd8135b4eb/test.json" | gojson -name=Speakers
type Speakers struct {
	Speaker []struct {
		Bio         string `json:"bio"`
		Category    string `json:"category"`
		Description string `json:"description"`
		Firstname   string `json:"firstname"`
		Image       string `json:"image"`
		Lastname    string `json:"lastname"`
		Link        string `json:"link"`
		Title       string `json:"title"`
	} `json:"speakers"`
}


func main() {

	resp, err := http.Get(url)
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
	defer resp.Body.Close()

	var speakers Speakers
	json.Unmarshal(bytes, &speakers)

	fmt.Println(len(speakers.Speaker))

	for _, speaker := range speakers.Speaker {
		fmt.Printf("%s-%s\n", speaker.Firstname, speaker.Lastname)
	}

	//testMarshal()


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

//func testMarshal() {
//	var speakers Speaker
//	speaker := Speaker{
//		"kevein",
//		"Leo",
//		"dev",
//		"rd",
//		"http://test.jsp",
//		"http://test.jsp",
//		"",
//		"",
//	}
//	speakers.Speakers = []Speaker{speaker}
//	result, _ := json.Marshal(speakers)
//	fmt.Println(string(result))
//}
