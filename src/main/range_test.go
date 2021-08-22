package main

import (
	"fmt"
	"testing"
)

func TestRange(t *testing.T) {

	// for循环的range格式
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// 可以通过赋值给 _ 来忽略序号和值。
	fmt.Println("only value")
	for _, v := range pow {
		fmt.Printf("%d\n", v)
	}

	// 如果只需要索引值，去掉 “ , value ” 的部分即可。
	fmt.Println("only index")
	for i := range pow {
		fmt.Printf("%d\n", i)
	}
}