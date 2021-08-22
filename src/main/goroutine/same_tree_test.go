package goroutine

import (
	"golang.org/x/tour/tree"
	"fmt"
	"testing"
)

func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

/**
	中序遍历两棵树，然后将遍历的结果分别放到channel中
	比较两个channel中的元素是否相同，即可判断是否是相同的树
 */
func Same(t1 *tree.Tree, t2 *tree.Tree) bool {
	if t1 == nil || t2 == nil {
		return t1 == t2
	}

	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	Walker(t1, ch1)
	Walker(t2, ch2)

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

func TestSameTree(test *testing.T) {
	ch := make(chan int, 100)
	t := tree.New(100)
	go Walker(t, ch)
	for v := range ch {
		fmt.Println(v)
	}

	if Same(tree.New(90), tree.New(90)) {
		fmt.Println("same tree")
	} else {
		fmt.Println("not the same")
	}

}

/**
	raw return，返回一个函数
 */
func Walker(t *tree.Tree, ch chan int) {
	func() {
		defer close(ch)
		Walk(t, ch)
	}()
}
