package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

type vertex struct {
	Lat, Long float64
}

/**
map 在使用之前必须用 make 来创建；值为 nil 的 map 是空的，并且不能对其赋值。
 */
var m map[string]vertex

var m2 = map[string]vertex{
	"Bell Labs": vertex{
		40.68433, -74.39967,
	},
	"Google": vertex{
		37.42202, -122.08408,
	},
}

/**
若顶级类型只是一个类型名，你可以在文法的元素中省略它
 */

var m3 = map[string]vertex{
	"Bell Labs": {
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
}

func main() {
	m = make(map[string]vertex)
	m["Bell Labs"] = vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
	fmt.Println(m2["Google"])
	fmt.Println(m3["Google"])

	// 修改map
	answerMap := make(map[string]int)

	answerMap["Answer"] = 42
	fmt.Println("The value: ", answerMap["Answer"])

	// 修改
	answerMap["Answer"] = 48
	fmt.Println("The value: ", answerMap["Answer"])

	// 删除
	delete(answerMap, "Answer")
	fmt.Println("The value: ", answerMap["Answer"])

	/**
	通过双赋值检测某个键存在：

elem, ok = m[key]
如果 key 在 m 中， ok 为 true。否则， ok 为 false，并且 elem 是 map 的元素类型的零值。

同样的，当从 map 中读取某个不存在的键时，结果是 map 的元素类型的零值。
	 */
	v, ok := answerMap["Answer"]
	fmt.Println("The value:", v, "Present? ", ok)

	// 词频统计
	wc.Test(WordCount)
}


func WordCount(s string) map[string]int {
	fields := strings.Fields(s)
	countMap := make(map[string]int)
	for i := 0; i < len(fields); i++ {
		word := fields[i]
		value, exists := countMap[word]
		if(!(exists)) {
			countMap[word] = 1
		} else {
			countMap[word] = value + 1
		}
	}
	return countMap
}