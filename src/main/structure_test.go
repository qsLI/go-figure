package main

import (
	"fmt"
	"testing"
)

type Vertex struct {
	X int
	Y int
}

func TestStruct(t *testing.T) {

	v1 := Vertex{1, 2}
	v1.X = 4
	p := &v1
	p.Y = 3
	// {4 3}
	fmt.Println(v1)
	// {1 0}
	v2 := Vertex{X: 1} // Y:0
	// {0 0}
	v3 := Vertex{}
	fmt.Println(v2)
	fmt.Println(v3)

	// 数组
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

}

