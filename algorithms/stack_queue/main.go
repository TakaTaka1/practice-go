package main

import (
	"container/list"
	"fmt"
)

func stackTest() {
	s := list.New()
	fmt.Println("--stack--\nPush:")
	for i := 0; i < 10; i++ {
		s.PushFront(i)
		fmt.Print(i, ",")
	}
	fmt.Println("\nPOP :")
	for i := 0; i < 10; i++ {
		r := s.Remove(s.Front())
		fmt.Print(r, ",")
	}
}

func queueTest() {
	s := list.New()
	fmt.Print("\n--Queue--\nEnqueue: ")
	for i := 0; i < 10; i++ {
		s.PushFront(i)
		fmt.Print(i, ",")
	}
	fmt.Print("\nDequeue: ")
	for i := 0; i < 10; i++ {
		r := s.Remove(s.Front())
		fmt.Print(r, ",")
	}
}

func main() {
	stackTest()
	queueTest()
}
