package main

import "fmt"

type Person struct {
	Name string
}

type XiaoMing struct {
	Person
	job string
}

func (p Person) walk()  {
	fmt.Println("walking")
}

func main() {
	p := Person{"unknown"}
	xm := XiaoMing{p, "teacher"}
	fmt.Println(xm.Person.Name)
	fmt.Println(xm.Name)
	xm.walk()
}
