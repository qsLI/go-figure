package main

import (
	"testing"
	"fmt"
	"time"
	"math/rand"
	"math"
)


func swap(s1, s2 string) (string, string) {
	return s2, s1
}

func TestSwap(t *testing.T) {

	fmt.Printf("hello, world\n")
	fmt.Printf("time is %s\n", time.Now())
	fmt.Printf("random number %d\n", rand.Intn(10))
	fmt.Printf("sqrt of 4 %f\n", math.Sqrt(4))
	fmt.Printf("value of PI %f\n", math.Pi)

	a, b := swap("Hello", "world")
	fmt.Println(a + " " + b)
}