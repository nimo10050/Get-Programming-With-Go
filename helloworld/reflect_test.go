package main

import (
	"reflect"
	"testing"
)

type Interface_0 interface {
	Func0()
}

type Struct_0 struct {
}

func (ss Struct_0) Func0() {
	println("ss func_0")
}

func (ss Struct_0) Func1() {
	println("ss func_1")
}

func (ss *Struct_0) Func2() {
	println("ss func_2")
}

func TestReflect(t *testing.T) {
	ss := Struct_0{}
	ss.Func0()
	ss.Func2()
	typ := reflect.TypeOf(ss)
	typ1 := reflect.TypeOf(&ss)
	//println(typ.Method(0))
	println("method num: ", typ.NumMethod())
	println("method num: ", typ1.NumMethod())
	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i)
		println(method.Name)
		//println("test")
		//println("method index : ", i, " method name: ", typ.Method(i))
	}
}
