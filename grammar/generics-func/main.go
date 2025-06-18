package main

import "fmt"

func add[T int | float64](a, b T) T {
	return a + b
}

func main() {
	num1 := add(1, 2)
	num2 := add(1.1, 2.2)
	fmt.Println(num1, num2)
}