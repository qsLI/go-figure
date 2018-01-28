package main

import (
	"fmt"
	"sync"
)

func int() {
	fmt.Println("initing...")
}

func main() {
	/**
		wg 用来等待程序完成
		类似java中的CountDownLatch
	  */
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("start goroutines")

	go func() {
		defer wg.Done()

		for count := 0; count < 100; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}

	}()

	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("all is done!")
}
