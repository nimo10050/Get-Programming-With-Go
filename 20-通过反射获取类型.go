package main

import (
	"fmt"
	"reflect"
)

type MyInterface interface {
	sayHello()
}

type Service1 struct {
}

type Service2 struct {
}

func (Service1) sayHello() {
	fmt.Println("hi, service1")
}

func (Service2) sayHello() {
	fmt.Println("hi, service2")
}

func main() {
	m := make(map[string]interface{}, 2)
	m["service1"] = Service1{}
	m["service2"] = Service2{}

	// 打印结果 main.Service1
	t := reflect.TypeOf(m["service1"])
	fmt.Println("service1 type is: ", t)

	// 打印结果 struct
	k := t.Kind()
	fmt.Println(k)
}
