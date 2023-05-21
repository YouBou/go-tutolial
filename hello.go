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
	fmt.Print("1から" + x + "の合計は、")
	t := 0
	c := 1
	for c <= n {
		t += c
		c++
	}
	fmt.Println(t, "です。")
}
