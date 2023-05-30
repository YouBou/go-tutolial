package main

import (
	"fmt"
)

func main() {
	// 整数値の代入
	n := 123
	// nのポインタを代入
	p := &n
	fmt.Println("number:", n)
	fmt.Println("pointer:", p) // アドレスの出力
	fmt.Println("value:", *p)  // ポインタにある値の出力
}
