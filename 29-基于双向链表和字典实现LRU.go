package main

import (
	"container/list"
	"fmt"
)

var (
	ll = list.New()
	m = make(map[string]*list.Element)
	m1 = make(map[string]list.Element)
)

func main()  {

	key := "key1"
	val := "value1"

	Add(key, val)

	fmt.Println("===容器初始值===")

	fmt.Println("map 中 key[", key, "]=", m[key].Value)
	fmt.Println("map1 中 key[", key, "]=", m1[key].Value)
	fmt.Println("list 中： ", ll.Front().Value)

	newVal := "value2"

	// 修改队列的 value
	Edit(newVal)

	fmt.Println("===修改容器后===")

	fmt.Println("map 中 key[", key, "]=", m[key].Value)
	fmt.Println("map1 中 key[", key, "]=", m1[key].Value)
	fmt.Println("list 中： ", ll.Front().Value)
}

func Edit(newVal string) {
	element := ll.Front()
	element.Value = newVal
}

func Add(key string, val interface{})  {
	element := ll.PushFront(val)
	m[key] = element
	m1[key] = *element
}