package main

import "fmt"

type stack struct {
	index	int
	data	[5]int
}

func (s *stack) push(k int) {
	s.data[s.index] = k
	s.index++
}

func (s *stack) pop() int {
	s.index--
	return  s.data[s.index]
}

func main() {
	// var address *int 	// declare an int pointer
	number := 42	 	// int
	address := &number	// address stores the memory address of number
	value := *address 	// dereferencung the value

	fmt.Printf("address: %v\n", address)
	fmt.Printf("value: %v\n", value)
}

