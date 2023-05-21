package main

import (
	"fmt"
	"strconv"

	"github.com/YouBou/go-tutolial/hello"
)

func main() {
	t := 0
	x := hello.Input("type a number")
	n, err := strconv.Atoi(x)
	if err != nil {
		goto err
	}
	for i := 1; i <= n; i++ {
		t += 1
	}
	fmt.Println("total:", t)
	return

err:
	fmt.Println("ERROR!")
}
