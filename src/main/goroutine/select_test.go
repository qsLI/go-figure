package goroutine

import (
	"fmt"
	"time"
	"testing"
)

/**
	select 语句使得一个 goroutine 在多个通讯操作上等待。
	select 会阻塞，直到条件分支中的某个可以继续执行，
	这时就会执行那个条件分支。
	当多个都准备好的时候，会随机选择一个。
 */
func fibonaccii(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <- quit:
			fmt.Println("quit")
			return
		}
	}
}


func TestSelectChannel(t *testing.T) {

	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<- c)
		}
		quit <- 0
	}()
	fibonaccii(c, quit)


	/**
	当 select 中的其他条件分支都没有准备好的时候，default 分支会被执行。
	为了非阻塞的发送或者接收，可使用 default 分支：
	 */
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}