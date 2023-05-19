package main

import (
	"fmt"
	"strconv"

	"github.com/YouBou/go-tutolial/hello"
)

func main() {
	x := hello.Input("type a number")
<<<<<<< HEAD
	fmt.Print(x + "月は、")
	switch n, err := strconv.Atoi(x); n {
	case 0:
		fmt.Println("整数値が得られません")
		fmt.Println(err)
	case 1, 2, 12:
		fmt.Println("冬です。")
	case 3, 4, 5:
		fmt.Println("春です。")
	case 6, 7, 8:
		fmt.Println("夏です。")
	case 9, 10, 11:
		fmt.Println("秋です。")
	default:
		fmt.Println("月の値ではありませんよ？")
||||||| bafa167
	fmt.Print(x + "は、")
	if n, err := strconv.Atoi(x); err == nil {
		if n%2 == 0 {
			fmt.Println("偶数です。")
			return
		}
		fmt.Println("奇数です。")
		return
=======
	n, err := strconv.Atoi(x)
	if err != nil {
		return
>>>>>>> master
	}
<<<<<<< HEAD
||||||| bafa167
	fmt.Println("整数ではありません。")
=======
	fmt.Print(x + "は、")
	switch {
	case n%2 == 0:
		fmt.Println("偶数です。")
	default:
		fmt.Println("奇数です。")
	}
>>>>>>> master
}
