package main


import (
	"golang.org/x/tour/pic"
	"math"
	"fmt"
	"strings"
)

func main() {
	/**
	[]T 是一个元素类型为 T 的 slice。
	 */
	s := []int{2, 3, 5, 7, 11, 13}
	var slice2 []int = s
	fmt.Println("s == ", s)
	fmt.Println(slice2)

	/**
	len(s) 返回 slice s 的长度。
	 */
	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] == %d\n", i, s[i])
	}


	// Create a tic-tac-toe board.
	game := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	game[0][0] = "X"
	game[2][2] = "O"
	game[2][0] = "X"
	game[1][0] = "O"
	game[0][2] = "X"

	printBoard(game)

	// 切片
	ss := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("ss[1:4] == ", ss[1:4])
	fmt.Println("ss[:3] == ", ss[:3])
	fmt.Println("ss[4:] == ", ss[4:])

	// 构造slice
	aa := make([]int, 5)
	printSlice("a", aa)
	bb := make([]int, 0, 5)
	printSlice("bb", bb)

	// slice的零值
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}

	// append works on nil slices
	z = append(z, 0)
	printSlice("z", z)

	// 扩容
	z = append(z, 2, 3, 4)
	printSlice("z", z)

	// for循环的range格式
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// 可以通过赋值给 _ 来忽略序号和值。
	for _, v := range pow {
		fmt.Printf("%d\n", v)
	}

	// 如果只需要索引值，去掉 “ , value ” 的部分即可。
	for i := range pow {
		fmt.Printf("%d\n", i)
	}

	smallPic := Pic(4, 4)
	for i, v := range smallPic {
		printSlice2(string(i), v)
	}

	pic.Show(Pic)

}

func Pic(dx, dy int) [][]uint8 {
	img := make([][]uint8, dy)
	for i, row := range img {
		for j := 0; j < dx; j++ {
			row = append(row, uint8(math.Pow(float64(j), 2)))
		}
		img[i] = row
	}
	return img
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func printSlice2(s string, x []uint8) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func printBoard(s [][]string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%s\n", strings.Join(s[i], " "))
	}
}