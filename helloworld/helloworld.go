package main

import (
	"container/list"
	"fmt"
)

/**
 * 定义一个 学生 结构体
 * == 感觉回到了 C 语言
 */
type Student struct {
	name      string
	age       int
	sex       byte
	studentNo string
}

/**
 * 定义存储学生的列表容器
 */
var studentList = list.New()

/**
 * 添加学生
 */
func add(student Student) {
	studentList.PushBack(student)
}

/**
 * 根据学号删除学生信息
 */
func remove(studentNo string) {
	// 打印变量类型
	// fmt.Println(reflect.TypeOf(student))
	// 从列表中返回的类型不是 Student 而是 *List.element
	for student := studentList.Front(); student != nil; student = student.Next() {
		// cast 返回的是否类型转换成功
		s, cast := (student.Value).(Student)
		if cast && studentNo == s.studentNo {
			fmt.Println(student)
			studentList.Remove(student)
			break
		}
	}
}

/**
 * 通过学号获取学生信息
 */
func get(studentNo string) Student {
	for student := studentList.Front(); student != nil; student = student.Next() {
		s, cast := (student.Value).(Student)
		if cast && studentNo == s.studentNo {
			return s
		}
	}
	// 无语子, 不知道怎么返回类似 java 中的 null
	return Student{}
}

/**
 * 编辑学生信息
 */
func update(updateStudent Student) {
	for student := studentList.Front(); student != nil; student = student.Next() {
		s, cast := (student.Value).(Student)
		if cast && updateStudent.studentNo == s.studentNo {
			// 先插入要更新的, 再删除以前的
			studentList.InsertAfter(updateStudent, student)
			studentList.Remove(student)
			break
		}
	}
}

/**
 * 打印学生信息
 */
func printStudent(student Student) {
	if student == (Student{}) {
		fmt.Println("Student is empty")
		return
	}
	// fmt.Println(reflect.DeepEqual(empty, student))
	var sex = ""
	if student.sex == 1 {
		sex = "男"
	} else {
		sex = "女"
	}
	fmt.Printf("学号: %s\n学生姓名: %s\n学生年龄: %d\n学生性别: %s\n", student.studentNo, student.name, student.age, sex)
}

/**
 * 程序入口
 */
func main() {

	// 添加一个学生信息
	student := Student{"zhangsan", 10, 1, "123"}
	add(student)

	// 打印列表长度
	// fmt.Printf("studentList.Len(): %v\n", studentList.Len())

	// 根据学号删除学生
	// remove("123")

	// 修改学生信息
	// student.age = 11
	//student.name = "lisi"
	//update(student)

	// 查询学生信息
	//var s = get("123")
	//fmt.Println(s)
	printStudent(student)
}
