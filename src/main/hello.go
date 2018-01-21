package main

import "fmt"
import "time"

import (
	"math/rand"
	"math"
	"math/cmplx"
	"runtime"
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
	fmt.Printf("hello, world\n")
	fmt.Printf("time is %s\n", time.Now())
	fmt.Printf("random number %d\n", rand.Intn(10))
	fmt.Printf("sqrt of 4 %f\n", math.Sqrt(4))
	fmt.Printf("value of PI %f", math.Pi)

	fmt.Printf("sum of 4 + 3 is %d\n", add(4, 3))

	a, b := swap("Hello", "world")
	fmt.Println(a + " " + b)

	fmt.Println(split(17))

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

	fmt.Printf("sum of 1 to 100 = %d\n", sum(100))

	result, e := sqrt(4)
	fmt.Printf("result: sqrt of %d is %d, error=%s\n", 4, result, e)
	result, e = sqrt(-4)
	fmt.Printf("result: sqrt of %d is %d, error=%s\n", -4, result, e)

	root := Sqrt4Newton(9.2)
	fmt.Printf("root of 9.2 is %f", root)

	// dead loop
	//for{
	//	fmt.Printf("in loop %s\n", time.Now())
	//	time.Sleep(70000000)
	//}

	// no more breaks~
	fmt.Println("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux. ")
	default:
		fmt.Printf("%s.", os)
	}

	/**
	没有条件的 switch 同 switch true 一样。

这一构造使得可以用更清晰的形式来编写长的 if-then-else 链。
	 */
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	/**
	defer 语句会延迟函数的执行直到上层函数返回。

延迟调用的参数会立刻生成，但是在上层函数返回前函数都不会被调用。
	 */
	defer fmt.Println("world")
	fmt.Println("hello")

	reverseCount(50)
}

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

func swap(s1, s2 string)(string, string) {
	return s2, s1
}

/**
Go 的返回值可以被命名，并且就像在函数体开头声明的变量那样使用。

返回值的名称应当具有一定的意义，可以作为文档使用。

没有参数的 return 语句返回各个返回变量的当前值。这种用法被称作“裸”返回。

直接返回语句仅应当用在像下面这样的短函数中。在长的函数中它们会影响代码的可读性。
 */
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

/**
Go 只有一种循环结构—— for 循环。

基本的 for 循环包含三个由分号分开的组成部分：

初始化语句：在第一次循环执行前被执行
循环条件表达式：每轮迭代开始前被求值
后置语句：每轮迭代后被执行
初始化语句一般是一个短变量声明，这里声明的变量仅在整个 for 循环语句可见。

如果条件表达式的值变为 false，那么迭代将终止。

注意：不像 C，Java，或者 Javascript 等其他语言，for 语句的三个组成部分 并不需要用括号括起来，但循环体必须用 { } 括起来。
顺序求和, 1..n
 */

 func sum(n int) int {
 	sum := 0
 	for i := 1; i <= n; i++ {
 		sum += i
	}
	return sum
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
	for ; z < x; z = z - (z*z - x)/(2*z) {
		fmt.Printf("iteration z = %f\n",  z)
		if math.Abs(z*z - x) < precision {
			break
		}
	}
	return z
}

func reverseCount(end int) {
	if end < 0 {
		return
	}

	for i := 0; i < end; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("reverse count done!")
}