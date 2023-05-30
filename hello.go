package main

import (
	"fmt"
)

func main() {
	ar := []int{10, 20, 30}
	fmt.Println(ar)
	initial(&ar)
	fmt.Println(ar)
}

func initial(ar *[]int) {
	for i := 0; i < len(*ar); i++ {
		// スライスの指定したインデックスのポインタではなくスライス自身のポインタを参照する
		(*ar)[i] = 0
	}
}
