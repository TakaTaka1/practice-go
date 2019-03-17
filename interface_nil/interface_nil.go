package main

import "fmt"

func main() {

	var (
		x, y interface{}
	)
	y = 2
	fmt.Print(x + y) // error

}
