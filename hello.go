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
	for i := 1; i <= n; i++ {
		t += i
	}
	fmt.Println(t, "です。")
}
