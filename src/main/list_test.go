package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()

	for i := 0; i < 100; i++ {
		l.PushBack(i)
	}

	fmt.Println(l.Front())
}