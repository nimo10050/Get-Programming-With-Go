package main

import "fmt"

type MyInterface interface {
	sayHello()
}

type MyStruct struct {
}

func (s *MyStruct) sayHello() {
	fmt.Println("hello everyone")
}

func main() {
	var ss = MyStruct{}
	ss2 := &ss
	ss.sayHello()
	ss2.sayHello()
}
