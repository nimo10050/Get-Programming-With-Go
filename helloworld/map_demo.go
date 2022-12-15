package main

import "fmt"

func main() {
	// 声明和初始化
	// key string, value int
	dict := make(map[string]int)
	dict["zhangsan"] = 41

	fmt.Println(dict["zhangsan"])
	// 方式二
	dict1 := map[string]int{"lisi": 1}
	fmt.Println(dict1["lisi"])

	// 判断 key 是否存在
	age, exists := dict["zhangsan"]

	fmt.Println("age: ", age, "exists: ",  exists)

	// delete key
	delete(dict, "zhangsan")

	age, exists2 := dict["zhangsan"]
	fmt.Println("age: ", age, "exists: ",  exists2)

	// 遍历 map
	for key, value := range dict1{
		fmt.Println("key: ", key, "value: ", value)
	}

	// 函数间传递 map


}
