package main

import (
	"fmt"
	"strconv"

	"github.com/YouBou/go-tutolial/hello"
)

func main() {
	x := hello.Input("type a number")
	fmt.Print(x + "は、")
	if n, err := strconv.Atoi(x); err == nil {
		if n%2 == 0 {
			fmt.Println("偶数です。")
			return
		}
		fmt.Println("奇数です。")
		return
	}
	fmt.Println("整数ではありません。")
}
