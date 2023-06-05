package main

import (
	"fmt"
)

// Data is interface.
type Data interface {
	// Data型interfaceを満たすメソッドの定義
	Initial(name string, data []int)
	PrintData()
}

// Mydata is Struct.
type Mydata struct {
	// Mydata型の構造体が持ちうる要素の定義
	Name string
	Data []int
}

// Initial is init method.
func (md *Mydata) Initial(name string, data []int) {
	// Mydata型の構造体の要素を初期化
	md.Name = name
	md.Data = data
}

// PrintData is println all data.
func (md *Mydata) PrintData() {
	// Mydata型構造体が持つ要素の出力
	fmt.Println("Name: ", md.Name)
	fmt.Println("Data: ", md.Data)
}

// Check is method.
func (md *Mydata) check() {
	fmt.Printf("Check! [%s]", md.Name)
}

func main() {
	// Data型の変数宣言
	var ob Mydata = Mydata{}
	// 構造体の値を定義
	ob.Initial("Sachiko", []int{55, 66, 77})
	// 構造体の要素を出力
	ob.check()
}
