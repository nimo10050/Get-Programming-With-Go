package main

import "fmt"

type Person interface {
	sayHello()
}

type Lisi struct {
	name string
}

func (lisi *Lisi) sayHello() {
	fmt.Println("hi, ", lisi.name)
}

// 接口转为 struct
// 空接口
var _ Person = (*Lisi)(nil)

func main() {
	var p Person = &Lisi{name: "zhangsan"}
	stu := p.(*Lisi) // 接口转为实例
	stu.sayHello()
	fmt.Printf("转换后的类型为: %T", stu)
}
