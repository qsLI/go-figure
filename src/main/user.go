package main

import "fmt"


/**
大写开头的才是对外暴露的, 序列化才会有
 */
type User struct {
	Name string `json:"name"`
	Age  int `json:"age"`
	Desc []string
}

func (user User) Do(in string) (string, int) {
	fmt.Println("%s Name is %s, Age is %d\n", in, user.Name, user.Age)
	return user.Name, user.Age
}