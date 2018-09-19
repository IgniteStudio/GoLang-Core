package main

import "fmt"

func main(){
// MAIN TYPES
// string
// bool
// int
// int int8 int16 int32 int64
// uint uint8 uint16 uint32 uint64 uintptr
// byte - alias for uint8
// rune - alias for int32
// float32 float64
// complex64 complex128

// Using var
//var name = "Dave"
// Short hand of above
name:= "Dave"
size:= 1.3
var age int32 = 37
var isCool = true


fmt.Println(name, age, isCool, size)
fmt.Printf("%T\n", isCool)

}