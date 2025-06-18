package main

import "fmt"

func main() {
	m1 := make(map[string]int)
	m1["age"] = 18
	fmt.Println("age",m1["age"])
	age,ok := m1["age"]
	fmt.Println("result",age,ok)
}
