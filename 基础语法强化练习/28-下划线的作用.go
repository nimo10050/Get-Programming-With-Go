package main

// 当使用 import 关键字引入包时， 如果我们只是想执行一下 go 文件里面的 init 方法， 而不是引用包里面的 go 文件
// 此时就能使用 如下方式：

import (
	"fmt"
	_ "io/ioutil"

)


func main()  {

	// 通过 var 关键字 实例化类型时， 下划线的作用是:
	var s = Student{Name: "sddss"}
	fmt.Println(s.GetUser())

	// 因为我们使用导入包时， 使用了 _, 所以我们是不能使用如下方式引用 go 文件中的内容的。
	// ioutil.ReadAll()


	// 当我们调用某个函数时， 这个函数有返回值， 但是我们并不需要过多 care 它时， 此时下划线的作用就跟占位符差不多
	_, _ = Quy(1)

	_, isOK := Quy(2)

	if isOK {
		fmt.Println("isOK: ", isOK)
	}

	err, isOK2 := Quy(3)

	if !isOK2 {
		fmt.Println("isOk2: ", isOK2, err)
	}

	// 通过 var _ 接口类型 = 对象 的方式， 定义变量，主要有两个目的：
	// 1. 编译期校验类型
	// 2. 如果变量不使用的话， 就用下划线当作占位符
	// https://www.modb.pro/db/390843
}

func Quy(i int) (error, bool)  {
	if i % 2 == 0 {
		return nil, true
	}
	err := fmt.Errorf("fuck you")
	return err, false
}

type User interface {
	GetUser() string
}


