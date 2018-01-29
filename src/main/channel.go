package main

import (
	"math/rand"
	"time"
	"sync"
	"fmt"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var wwg sync.WaitGroup

func main() {

	court := make(chan int)

	wwg.Add(2)

	go player("Nadal", court)
	go player("Djokovic", court)

	court <- 1

	wwg.Wait()

}

func player(name string, court chan int) {
	defer wwg.Done()

	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("Player %s Won\n", name)
			return
		}

		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)
			close(court)
			return
		}

		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		court <- ball
	}
}