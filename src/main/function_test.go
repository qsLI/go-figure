package main

import (
	"math"
	"fmt"
	"time"
	"strings"
	"io"
)

import (
	"golang.org/x/tour/reader"
	"os"
	"net/http"
	"log"
	"testing"
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

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

/**
类型通过实现那些方法来实现接口。 没有显式声明的必要；所以也就没有关键字“implements“。
隐式接口解藕了实现接口的包和定义接口的包：互不依赖。
 */
type Abser interface {
	Abs() float64
}

type Person struct {
	Name string
	Age int
}

type IPAddr [4]byte

func (addr IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", addr[0], addr[1], addr[2], addr[3])
}

/**
Stringer 是一个可以用字符串描述自己的类型。`fmt`包 （还有许多其他包）使用这个来进行输出。
 */
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

type MyError struct {
	When time.Time
	What string
}

/**
与 fmt.Stringer 类似， error 类型是一个内建接口：

type error interface {
    Error() string
}
（与 fmt.Stringer 类似，fmt 包在输出时也会试图匹配 error。）


 */
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func TestFunc(t *testing.T) {
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

	myFloat := MyFloat(-math.Sqrt2)
	fmt.Println(myFloat.Abs())

	var a Abser
	a = myFloat
	// a = *v // not assignable
	fmt.Println(a.Abs())

	p1 := Person{"Andrew", 45}
	p2 := Person{"Zippo", 32}

	fmt.Println(p1, p2)

	var ip IPAddr = [4]byte {8, 8, 8, 8}
	fmt.Println(ip)

	if err := run(); err != nil {
		fmt.Println(err)
	}


	r := strings.NewReader("Hello, Reader!")
	b := make([]byte, 8)

	for {
		n, e := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, e, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if e == io.EOF {
			break
		}
	}

	reader.Validate(MyReader{})

	// rot13
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	rot13Reader := rot13Reader{s}
	io.Copy(os.Stdout, &rot13Reader)
	fmt.Println("")

	//var h Hello
	//err := http.ListenAndServe("127.0.0.1:4000", h)
	//if err != nil {
	//	log.Fatal(err)
	//}
	fmt.Println("serve at 127.0.0.1:4000, endpoint /struct")
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	err := http.ListenAndServe("127.0.0.1:4000", nil)
	if err != nil {
		log.Fatal(err)
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

type MyReader struct {}

func (r MyReader) Read(b []byte) (int, error) {
	for i := 0; i < len(b); i++ {
		b[i] = byte('A')
	}
	return len(b), nil
}

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	count, e := r.r.Read(b)
	for i, v := range b[:count] {
		if v >= 'A' && v <= 'Z' {
			b[i] = (v + 13 - 'A') % 26 + 'A'
		} else  if v >= 'a' && v <= 'z' {
			b[i] = (v + 13 - 'a') % 26 + 'a'
		}
	}
	return count, e
}

type Hello struct {}

/**
包 http 通过任何实现了 http.Handler 的值来响应 HTTP 请求：

package http

type Handler interface {
    ServeHTTP(w ResponseWriter, r *Request)
}
 */
func (h Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
		fmt.Fprint(w, "Hello!")
}


type Struct struct {
	Greeting string
	Punct 	 string
	Who      string
}


func (s Struct) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, fmt.Sprintf("%s %s %s", s.Who, s.Punct, s.Greeting))
}