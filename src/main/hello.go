package main

import "fmt"

import (
	"math"
	"math/cmplx"
)

/**
var 语句定义了一个变量的列表；跟函数的参数列表一样，类型在后面。

就像在这个例子中看到的一样， var 语句可以定义在包或函数级别。
 */
var c, python, java bool = false, true, true

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
	Big   = 1 << 100
	Small = Big >> 99
)

const precision float64 = 0.00001

func main() {

	/**
	在函数中， := 简洁赋值语句在明确类型的地方，可以用于替代 var 定义。

函数外的每个语句都必须以关键字开始（ var 、 func 、等等）， := 结构不能使用在函数外。
	 */
	var i int
	k := 77
	fmt.Println(i, c, python, java, k)
	// go 没有指针运算
	var p *int = &k
	fmt.Println(p)
	fmt.Println(*p)

	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)

	floatNumber := float64(k)
	u := uint(k)

	fmt.Printf(f, floatNumber, floatNumber)
	fmt.Printf(f, u, u)

	const World = "世界"
	fmt.Println(World)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	result, e := sqrt(4)
	fmt.Printf("result: sqrt of %d is %d, error=%s\n", 4, result, e)
	result, e = sqrt(-4)
	fmt.Printf("result: sqrt of %d is %d, error=%s\n", -4, result, e)

	root := Sqrt4Newton(9.2)
	fmt.Printf("root of 9.2 is %f", root)
}

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

/**
if语句也不需要圆括号
 */
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", float64(e))
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return -1, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

/**
牛顿法求根: https://zh.wikipedia.org/wiki/File:NewtonIteration_Ani.gif
 */
func Sqrt4Newton(x float64) (float64) {
	if x < 0 {
		return -1
	}

	z := float64(x / 2)
	for ; z < x; z = z - (z*z-x)/(2*z) {
		fmt.Printf("iteration z = %f\n", z)
		if math.Abs(z*z-x) < precision {
			break
		}
	}
	return z
}
