package main

import (
	"fmt"
	"strconv"

	"github.com/YouBou/go-tutolial/hello"
)

func main() {
	const TAX_RATE = 1.1
	x := hello.Input("type a price")
	n, err := strconv.Atoi(x)
	if err != nil {
		fmt.Println("ERROR!!")
		return
	}
	p := float64(n)
	fmt.Println(int(p * TAX_RATE))
}
