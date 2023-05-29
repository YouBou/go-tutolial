package main

import (
	"fmt"
)

func main() {
	data := "*新しい値*"
	m1 := modify(data)
	data = "+new data+"
	m2 := modify(data)

	fmt.Println(m1())
	fmt.Println(m2())
}

// stringを引数に、スライス(string)を返り値とした関数を返す
func modify(d string) func() []string {
	// スライス(string)の定義
	m := []string{
		"1st", "2nd",
	}
	// 引数のstringを初期化したスライスに追加する関数
	return func() []string {
		return append(m, d)
	}
}
