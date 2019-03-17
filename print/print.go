package main

import "fmt"

func main() {
	// 文字列を出力
	fmt.Printf("数値=%v 文字列=%v  配列=%v\n", 5, "Golang", [...]int{1, 2, 3})

	// %#vはGoのリテラル表現でデータを埋め込む
	fmt.Printf("数値=%#v 文字列=%#v 配列=%#v", 5, "Golang", [...]int{1, 2, 3})

	// 型を出力
	fmt.Printf("数値=%#T 文字列%#T\n 配列=%#T", 5, "Golang", [...]int{1, 2, 3})

	println("Hello world")
	print("Hello world")
}
