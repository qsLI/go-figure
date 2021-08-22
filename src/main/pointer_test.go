package main

import (
	"testing"
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

/**
Go 没有类。然而，仍然可以在结构体类型上定义方法。
方法接收者 出现在 func 关键字和方法名之间的参数中。即 -> *Point
 */
func (v *Point) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

/**
刚刚看到的两个 Abs 方法。一个是在 *Vertex 指针类型上，而另一个在 MyFloat 值类型上。
有两个原因需要使用指针接收者。
	1. 首先避免在每个方法调用中拷贝值（如果值类型是大的结构体的话会更有效率）。
	2. 其次，方法可以修改接收者指向的值。
 */
func (v *Point) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

/**
not working
 */
func (v Point) Scale2(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}


func TestPointerModify(t *testing.T) {

	p := Point{3, 4}
	v := &p
	// 方法是定义在Pointer Reciever上的，这里p自动给转成了Pointer
	fmt.Println(p.Abs())
	fmt.Println(v.Abs())

	v.Scale(2)
	fmt.Println("scaled : ", v)

	// 没有生效，没有传址
	v.Scale2(2)
	fmt.Println("scaled2 : ", v)

}

// go 没有指针运算 p + 1这种
func TestPointer(t *testing.T) {
	i, j := 42, 2701

	p := &i         // point to i
	// 42
	fmt.Println(*p) // read i through the pointer
	// address
	fmt.Println(p)
	*p = 21         // set i through the pointer
	// 21
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	// 73
	fmt.Println(j) // see the new value of j
}