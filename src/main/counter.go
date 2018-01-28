package main

import (
	"sync"
	"fmt"
	"sync/atomic"
	"runtime"
)

var (
	counter int64
	wg sync.WaitGroup
)

func main() {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()

	fmt.Println("Final Counter: ", counter)
}

func incCounter(id int) {
	defer wg.Done()

	for count:=0; count < 2; count++ {
		/**
		原子操作
		 */
		atomic.AddInt64(&counter, 1)

		/**
		当前goroutine从线程推出, 并放回到队列
		 */
		runtime.Gosched()
	}
}