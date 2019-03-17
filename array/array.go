package main

import "fmt"

func main() {
	// array
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%v", a)

	// 要素数の省略
	b := [...]uint8{1, 2, 3}
	b[0] = 255
	fmt.Print(b) // OK

	// arrayの代入
	a1 := [3]int{1, 2, 3}
	a2 := [3]int{4, 5, 6}

	a1 = a2
	fmt.Printf("%v", a1)

	a1[0] = 0
	a1[1] = 0

	fmt.Printf("%v", a1)
	fmt.Printf("%v", a2)
}
