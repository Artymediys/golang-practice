package main

import "fmt"

func option1() {
	a, b := 42, 13
	fmt.Printf("a: %d, b: %d\n", a, b)

	a, b = b, a
	fmt.Printf("a: %d, b: %d\n", a, b)
}

func option2() {
	a, b := 42, 13 // a: 101010, b: 1101
	fmt.Printf("a: %d, b: %d\n", a, b)

	// Применяем XOR
	a ^= b // a: 100111, b: 1101
	b ^= a // a: 100111, b: 101010
	a ^= b // a: 1101, b: 101010
	fmt.Printf("a: %d, b: %d\n", a, b)
}

func option3() {
	a, b := 42, 13
	fmt.Printf("a: %d, b: %d\n", a, b)

	a += b    // a: 55, b: 13
	b = a - b // a: 55, b: 42
	a -= b    // a: 13, b: 42
	fmt.Printf("a: %d, b: %d\n", a, b)
}

func main() {
	option1()
	option2()
	option3()
}
