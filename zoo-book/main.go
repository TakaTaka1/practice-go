package main

import (
	"fmt"

	"./animals"
)

func main() {
	fmt.Println(AppName())

	fmt.Println(animals.MonkeyFeed())
	fmt.Println(animals.ElephantFeed())
	fmt.Println(animals.RabbitFeed())

	fmt.Printf("数値=%#v 文字列%#v\n 配列=%#v", 5, "Golang", [...]int{1, 2, 3})
	fmt.Printf("数値=%#T 文字列%#T\n 配列=%#T", 5, "Golang", [...]int{1, 2, 3})
}
