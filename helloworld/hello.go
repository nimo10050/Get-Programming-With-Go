package main

import (
	"container/list"
	"fmt"
	"reflect"
)

type Student struct {
	name      int
	age       int
	sex       int
	studentNo int
}

func testFanxing() *list.List {
	myList := list.New()
	myCar := Car{"xiaoch"}
	myList.PushBack(myCar)
	return myList
}

type Car struct {
	Name string `json:"name"`
}

func main() {
	fmt.Println("hello world")
	var i = 1
	switch i {
	case 0:
	case 1:
		fmt.Println(i)
	}

	myArray := []int{1, 2, 3}
	for index := range myArray {
		myArray[index] = myArray[index] * 2
	}

	for index := range myArray {
		fmt.Println(myArray[index])
	}

	fmt.Printf("%s   %+v\n", "a", "c")

	myMap := make(map[string]string, 10)
	myMap["key"] = "value"
	fmt.Println(myMap["key10"])

	//myStudent := Student{}
	//p := *myStudent
	//fmt.Println(&myStudent)
	//fmt.Println(*myStudent)
	myCar := Car{Name: "zhangsan"}
	myType := reflect.TypeOf(myCar)
	name := myType.Field(0)
	myTag := name.Tag.Get("json")
	fmt.Println(myTag)

}
