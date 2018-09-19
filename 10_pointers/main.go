package main

import "fmt"

func main(){
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	// Use * to read val from mem address
	fmt.Println(*b)
	fmt.Println(*&a) // NOTE: * and & cancel each other out 

	// Change val with pointer
	*b = 10
	fmt.Println(a)
}