package main

import "fmt"

var a uint8 = 1

var x1 uint8 = a << 1
var x2 uint8 = a << 5

var y1 uint8 = a << 1
var y2 uint8 = a << 2

var x uint8 = x1 | x2
var y uint8 = y1 | y2

func main() {
	fmt.Printf("1 = %08b\n", a)
	fmt.Printf("1 << 1 = %08b\n", x1)
	fmt.Printf("1 << 5 = %08b\n", x2)

	fmt.Printf("1 << 2 = %08b\n", y2)	
	
	fmt.Printf("1 << 1 | 1 << 5 = %08b\n", x)
	fmt.Printf("1 << 1 | 1 << 2 = %08b\n", y)

	fmt.Printf("x = %08b\n", x)
	fmt.Printf("y = %08b\n", y)
	fmt.Printf("x & y = %08b\n", x&y)
	fmt.Printf("x | y = %08b\n", x|y)
	fmt.Printf("x ^ y = %08b\n", x^y)
}