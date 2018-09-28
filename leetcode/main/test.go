package main

import "fmt"

func main() {
	array := make([]int, 10, 10)
	fmt.Println(cap(array))
	fmt.Println(len(array))

	s := array[2:5]
	fmt.Println(cap(s))
	fmt.Println(len(s))
}
