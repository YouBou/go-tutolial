package main

import (
	"fmt"
	"strconv"

	"github.com/YouBou/go-tutolial/hello"
)

func main() {
	x := hello.Input("type a number")
	n, err := strconv.Atoi(x)
	if err != nil {
		return
	}
	fmt.Print(x + "は、")
	switch {
	case n%2 == 0:
		fmt.Println("偶数です。")
	default:
		fmt.Println("奇数です。")
	}
}
