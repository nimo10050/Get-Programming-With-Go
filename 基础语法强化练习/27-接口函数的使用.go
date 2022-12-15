package main

import "fmt"

// Getter 定义一个接口
type Getter interface {
	Get(key string) string
}

type Func func(string) string

func (f Func) Get(key string) string {
	return f(key)
}

//func Load(key string, defaultVal string, f Func) string {
//	if key == "" && f != nil {
//		return f(key)
//	}
//	return defaultVal
//}

func Load1(key string, getter Getter) string {
	if key != "" && getter != nil {
		return getter.Get(key)
	}
	return "empty string"
}

type DB struct {
}

func (db DB) Get(key string) string {
	return "db"
}

func main() {
	// 相比于直接将函数对象作为 函数的入参来讲,
	// 接口函数, 既能够将普通的函数类型（需类型转换）作为参数，也可以将结构体作为参数，使用更为灵活，可读性也更好，这就是接口型函数的价值。
	// 类似于 Java8 的函数表达式
	val := Load1("1111", &DB{})

	fmt.Println("val: ", val)

	val1 := Load1("1111", Func(func(s string) string {
		fmt.Println("call back")
		return s
	}))

	fmt.Println("val1: ", val1)
}
