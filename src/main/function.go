package main

import (
	"math"
	"fmt"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}


func main() {
	hypot := func(x, y float64) float64{
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	/**
	  闭包,函数 adder 返回一个闭包。每个返回的闭包都被绑定到其各自的 sum 变量上。
	  */
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	// fibonacci
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

// fibonacci 函数会返回一个返回 int 的函数。
func fibonacci() func() int {
	var fn_1 int = 1
	var fn_2 int = 1
	return func() int {
		result := fn_1 + fn_2
		fn_2 = fn_1
		fn_1 = result
		return result
	}
}
