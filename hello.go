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
	for _, v := range ar {
		n, er := strconv.Atoi(v)
		if er != nil {
			error()
			return
		}
		t += n
	}
	fmt.Println("total:", t)
}

func error() {
	fmt.Println("ERROR!")
}
