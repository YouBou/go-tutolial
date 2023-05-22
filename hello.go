package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/YouBou/go-tutolial/hello"
)

func main() {
	x := hello.Input("input data")
	ar := strings.Split(x, ",")
	t := 0
	for i := 0; i < len(ar); i++ {
		n, er := strconv.Atoi(ar[i])
		if er != nil {
			error()
		}
		t += n
	}
	fmt.Println("total:", t)
}

func error() {
	fmt.Println("ERROR!")
}
