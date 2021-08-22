package main

import (
	"encoding/json"
	"log"
	"fmt"
)

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
