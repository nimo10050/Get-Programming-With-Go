package main

import (
	"fmt"
)

type Person1 struct {
	Name string
	Age  int
}

type ByAge []Person1

func MyFunc() {
	fmt.Println("MyFunc")
}

type MyFuncAlisa MyFunc

func main() {
	p := []Person1{
		{"wangwu", 2},
		{"zhangsan", 1},
		{"lisi", 3},
	}

	age := ByAge(p)
	fmt.Println(age)
	fmt.Printf("age type: %T\n", age)
	// sort.Sort()
	// fmt.Println(p)
}

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
