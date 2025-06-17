package main

import (
	"fmt"
	_ "grammar/base/test01"
	"grammar/base/test02"
)

func main() {
	test02.Add()
	fmt.Println("请输入你的年龄")
	var age uint8
	fmt.Scanln(&age)
	if age < 13 {
		fmt.Println("年龄小于13")
	} else {
		fmt.Println("年龄大于13")
	}
}
