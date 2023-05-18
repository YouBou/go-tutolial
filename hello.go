package main

import (
	"fmt"

	"github.com/YouBou/go-tutolial/hello"
)

func main() {
	name := hello.Input("type your name")
	fmt.Println("Hello, " + name + "!!")
}
