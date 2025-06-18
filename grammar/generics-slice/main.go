package main

import "fmt"

type CustomerSlice[T int | float64] []T

func main(){
	var mySlice1 CustomerSlice[int]
	var mySlice2 CustomerSlice[float64]
	mySlice1 = append(mySlice1, 1, 2, 3)
	mySlice2 = append(mySlice2, 1.1, 2.2, 3.3)
	fmt.Println(mySlice1) // [1 2 3]
	fmt.Println(mySlice2) // [1.1 2.2 3.3]
}