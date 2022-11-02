package main

import (
	"fmt"
	"reflect"
)

type MyInterface interface {
	sayHello()
}

type Service1 struct {
	Name string
	Age  int
	sex  int
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
	m["service1"] = Service1{Name: "zhangsan", Age: 1, sex: 1}
	m["service2"] = Service2{}

	// 拿到对象
	o := m["service1"]
	// 拿到类型
	t := reflect.TypeOf(o)
	fmt.Println("service1 type is: ", t)
	// 拿到值
	v := reflect.ValueOf(o)
	// 拿到字段
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Println("f.Name: ", f.Name, ", f.Type: ", f.Type, ", f.Value: ", v.Field(i))
	}
}
