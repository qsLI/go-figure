package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

func Same(t1 *tree.Tree, t2 *tree.Tree) bool {
	if t1 == nil || t2 == nil {
		return t1 == t2
	}

	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	Walk(t1, ch1)
	Walk(t2, ch2)
	close(ch1)
	close(ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if !ok1 || !ok2 {
			return ok1 == ok2
		}

		if v1 != v2 {
			break
		}
	}

	return false
}

func main() {
	ch := make(chan int, 100)
	t := tree.New(100)
	go func() {
		defer close(ch)
		Walk(t, ch)
	}()
	for v := range ch {
		fmt.Println(v)
	}

	if Same(tree.New(90), tree.New(90)) {
		fmt.Println("same tree")
	 } else {
	 	fmt.Println("not the same")
	}

}
