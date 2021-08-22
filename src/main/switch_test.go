package main

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestSwitch(t *testing.T) {

	// no more breaks~
	fmt.Println("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux. ")
	default:
		fmt.Printf("%s.", os)
	}
}

func TestSwitchWithoutVar(t *testing.T) {
	/**
		没有条件的 switch 同 switch true 一样。
		这一构造使得可以用更清晰的形式来编写长的 if-then-else 链。
 	*/
	now := time.Now()
	fmt.Println(now)
	switch {
	case now.Hour() < 12:
		fmt.Println("Good morning!")
	case now.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}