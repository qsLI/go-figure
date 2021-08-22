package main

import (
	"net/http"
	"log"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

const (
	baiduUrl = "http://ww.baidu.com/s?"
)

func main() {
	response, err := Get()

	document, err := goquery.NewDocumentFromResponse(response)
	if err != nil {
		log.Fatalln("goquery error!", err)
	}
	result := document.Find("#content_left .result").Text()
	fmt.Println(result)

	options := map[string]int{}
	options["汤素兰"] = 0
	options["叶圣陶"] = 0
	options["金波"] = 0
	options["郑渊洁"] = 0

	for key := range options {
		options[key] = strings.Count(result, key)
	}

	fmt.Println(options)
}
func Get() (*http.Response, error) {
	values := url.Values{}
	values.Add("wd", "80后经典动画片魔方大厦的原著作者是哪位童话作家")
	response, err := http.Get(baiduUrl + values.Encode())
	if err != nil {
		log.Fatalln("request error!", err)
	}
	defer response.Body.Close()
	log.Println(response.Header)
	return response, err
}
