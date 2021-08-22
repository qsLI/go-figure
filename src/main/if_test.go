package main

import (
	"testing"
	"fmt"
	"math"
)

func TestIF(t *testing.T) {
	if true {
		fmt.Println("just a test")
	}

	if a := 100; math.Pow(float64(a), 2) > 50 {
		fmt.Println("if with statement")
	}
}


func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func TestElse(t *testing.T) {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
