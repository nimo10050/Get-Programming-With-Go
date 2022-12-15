package main

import "fmt"

type Person interface {
	sayHello()
}

type Lisi struct {
	name string
}

// 注释掉这个方法
// 编译报错如下：
// Cannot use '&Lisi{name: "zhangsan"}' (type *Lisi) as the type Person
// Type does not implement 'Person' as some methods are missing: sayHello()
func (lisi *Lisi) sayHello() {
	fmt.Println("hi, ", lisi.name)
}

func main() {
	var p Person = &Lisi{name: "zhangsan"}
	p.sayHello()
	fmt.Printf("转换后的类型为: %T", p)
}
