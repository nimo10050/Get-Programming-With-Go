package main

import (
	"fmt"
	"reflect"
)

type MyInterface interface {
	sayHello()
}

type Service struct {
}

func (Service) SayHello() {
	fmt.Println("hi, service1")
}

func main() {
	o := Service{}

	// 拿到类型
	v := reflect.ValueOf(o)

	// 拿到方法
	for i := 0; i < v.NumMethod(); i++ {
		m := v.Method(i)
		m.Call(nil)
	}
}
