package main

import (
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {
	/**
		defer 语句会延迟函数的执行直到上层函数返回。
		延迟调用的参数会立刻生成，但是在上层函数返回前函数都不会被调用。
	 */
	var i int = 0
	defer fmt.Println("world", i)
	i = i + 1
	fmt.Println("hello", i)
}

func TestReverseCount(t *testing.T) {

	end := 100

	for i := 0; i < end; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("reverse count done!")
}