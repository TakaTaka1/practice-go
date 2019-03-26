package main

import "fmt"

func main() {
	// Your code here!
	a := [8]int{9, 4, 10, 3, 2, 7, 0, 1}
	var tmp = 0
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if a[i] < a[j] {
				tmp = a[j]
				a[j] = a[i]
				a[i] = tmp
			}
		}
	}

	fmt.Println(a)

}
