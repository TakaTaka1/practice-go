package main

import (
	"fmt"
	"math"
)

// パッケージ変数
var n = 100

func dosomething(a, b uint32) bool {
	//var sum uint32
	if (math.MaxUint32 - a) < b {
		return false
	} else {
		//sum = a + b
		return true
	}
}

func main() {
	//	tmp := dosomething(10, math.MaxUint32)
	//tmp2 := dosomething(10, math.MaxUint32)

	if dosomething(11, math.MaxUint32-10) == true {
		fmt.Printf("true\n")
	} else {
		fmt.Printf("false\n")
	}
	//一度にまとめてint型で
	//	var x, y, z int

	//	//変数ごとに型を分ける
	//	var (
	//		g, h int
	//		name string
	//	)

	// 型推論
	//	name := "test\n"
	//	fmt.Printf(name)

	//	n = n + 1
	//	fmt.Printf("n=%d", n)

	// 参照する変数が1つだとエラーになる
	//	var a, b, c int
	//	a = 1
	//	b = 1
	//	c = 1
	//	print(a)

	// 実装依存により下記はコンパイルエラーとなる
	//	var (
	//		n1 int
	//		n2 int64
	//	)

	//	n1 = 1
	//	n2 = n1
	//	fmt.Println(n2)

	// 下記は０が返り値となる()
	//	var by byte
	//	pass := 256
	//	by = byte(pass)
	//	fmt.Print(by)

	//  b3 = 254になる
	//	b1 := byte(255) //byte = uint8
	//	b3 := b1 + byte(255)
	//	fmt.Print(b3)

	// 下記のsumは44億は表示しない(ラップアラウンドをするため)
	//	var ui_1 = uint32(400000000)
	//	var ui_2 = uint32(4000000000)
	//	sum := ui_1 + ui_2
	//	fmt.Printf("%d + %d = %d", ui_1, ui_2, sum)
	//	fmt.Printf("uint32 max value=%d\n", uint32(4294967295)-1)
	//	fmt.Printf("uint32 max value=%d\n", math.MaxUint32)

	// 指数指定
	fmt.Print(1e3, "\n")

	// ルーンリテラル
	tmp := '\\' // シングルクォーテーションが必須
	tmp2 := 'i'
	fmt.Printf("%v%v\n", tmp, tmp2) // ダブルクォーテーションが必須
	//	fmt.Println(tmp)
	//	fmt.Print(tmp)

	// Rawリテラル バッククォート
	raw := `\n\n\n`
	fmt.Print(raw)

}
