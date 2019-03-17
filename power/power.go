package main

import "fmt"

func main() {
	var num int
	var pow int
	var result = 1

	num = 2
	pow = 4

	for i := 0; i < pow; i++ {
		result *= num
	}

	fmt.Printf("%dの%d乗は%dです。\n", num, pow, result)
}

// セミコロンは自動補完
// var a = 2; var b = 1 左記の場合は必要
// 型名を変数名の後ろに書く
// 宣言と同時に初期化を行う場合は型を省略できる
// var result =1
// 宣言と代入を同時に
// msg := "hello go"

// if a < 5 {} 丸括弧は不要 付けてもいいが冗長
// ブロック内が単文でも{}は必要
