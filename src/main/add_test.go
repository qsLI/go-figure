package main

import (
	"fmt"
	"testing"
)

/**
当两个或多个连续的函数命名参数是同一类型，则除了最后一个类型之外，其他都可以省略。

在这个例子中 ，

x int, y int
被缩写为

x, y int
 */
func add( x , y int) int {
	return x + y
}


func TestAdd(t *testing.T) {
	fmt.Printf("sum of 4 + 3 is %d\n", add(4, 3))
}