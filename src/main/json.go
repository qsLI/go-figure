package main

import (
	"encoding/json"
	"log"
	"fmt"
)

/**
大写开头的才是对外暴露的, 序列化才会有
 */
type User struct {
	Name string `json:"name"`
	Age  int `json:"age"`
	Desc []string
}

func main() {
	user := User{
		"kevin leo",
		77,
		[]string{"strong", "fast"},
	}

	fmt.Println(user.Name)

	bytes, e := json.Marshal(user)

	if e != nil {
		log.Fatal("ERROR: marshal error!")
	}

	fmt.Println(string(bytes))
}