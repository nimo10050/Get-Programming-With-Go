package main

import "fmt"

type myStruct struct {
	Name string
	Age  int
}

func main() {
	s := myStruct{}
	s.Name = "zhangsan"
	s.Age = 10
	fmt.Printf("%+v", s)
}
