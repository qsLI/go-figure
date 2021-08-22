package goroutine

import (
	"time"
	"fmt"
	"testing"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sam(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}

	c <- sum // ch <- v    // 将 v 送入 channel ch。
}

func fibonaccci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	/**
		只有发送者才能关闭 channel, 向一个已经关闭的 channel 发送数据会引起 panic
		通常情况下无需关闭它们。只有在需要告诉接收者没有更多的数据的时候才有必要进行关闭，
		例如中断一个 range
	  */
	close(c)
}

func TestGoRoutine(t *testing.T) {
	/**
		开启一个新的 goroutine 执行
	 */
	go say("world")
	say("hello")

	a := []int{7, 2, 8, -9, 4, 0}

	/**
		向带缓冲的 channel 发送数据的时候，只有在缓冲区满的时候才会阻塞。
		而当缓冲区为空的时候接收操作会阻塞。
	 */
	c := make(chan int, 100) // 带缓冲的

	/**
		默认情况下，在另一端准备好之前，发送和接收都会阻塞。
		这使得 goroutine 可以在没有明确的锁或竞态变量的情况下进行同步。
	 */
	go sam(a[:len(a)/2], c)
	go sam(a[len(a)/2:], c)
	//v := <-ch  // 从 ch 接收，并且赋值给 v。 （“箭头”就是数据流的方向。）
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)

	// buffered
	c2 := make(chan int, 10)
	go fibonaccci(cap(c2), c2)
	// 会不断从 channel 接收值，直到它被关闭。
	for i := range c2 {
		fmt.Println(i)
	}

	_, ok := <- c2
	if ok == false {
		fmt.Println("channel is closed")
	}
}