package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
  Name     string `json:"name"`
  Age      int    `json:"age,omitempty"`
  Password string `json:"_"`
}

func main() {
  s := Student{
    Name:     "枫枫",
	Age: 0,
    Password: "123456",
  }
  byteData, _ := json.Marshal(s)
  fmt.Println(string(byteData)) // {"name":"枫枫","age":21}
}