package main

import "fmt"

type Student[T int | float64] struct {
	Name string
	Age  int
	Score T
}

func main() {
	s1 := Student[int]{
		Name: "s1",
		Age:  18,
		Score: 100,
	}
	s2 := Student[float64]{
		Name: "s2",
		Age:  18,
		Score: 99.9,
	}
	fmt.Println(s1) // {s1 18 100}
	fmt.Println(s2) // {s2 18 99.9}
}