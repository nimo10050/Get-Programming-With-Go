package main

import "fmt"

type myInterface interface {
	sayHello()
}

type littleStruct struct {
}

func (littleStruct) sayHello() {
	fmt.Println("hi~")
}

func main() {
	s := littleStruct{}
	fmt.Printf("s 的类型： %T \n", s)
	s.sayHello()

	// 取 s 的引用
	ss := &s
	fmt.Printf("ss 的类型： %T \n", s)
	ss.sayHello()
}
