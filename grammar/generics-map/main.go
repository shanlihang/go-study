package main

import "fmt"

type MyMap[K string | int, V any] map[K]V

func main() {
	var myMap = make(MyMap[string, string])
	myMap["name"] = "test"
	fmt.Println(myMap) // map[name:test]
}