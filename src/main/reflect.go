package main

import (
	"reflect"
	"fmt"
)


/**
大写开头的才是对外暴露的, 序列化才会有
 */
type People struct {
	Name string `json:"name"`
	Age  int `json:"age"`
	Desc []string
}

func (user People) Do(in string) (string, int) {
	fmt.Println("%s Name is %s, Age is %d\n", in, user.Name, user.Age)
	return user.Name, user.Age
}

func main() {

	u := People{"tom", 6, nil}

	v := reflect.ValueOf(u)
	fmt.Println(v)

	t := reflect.TypeOf(u)
	fmt.Println(t)
}