package main

import (
	"container/list"
	"fmt"
)

var (
	l = list.New()
	cache = make(map[string]*list.Element)
)

type Val interface {
	Length() int64
}

type Entry struct {
	key string
	value Val
}

type SS string

func (ss SS) Length() int64 {
	return int64(len(ss))
}

// 这个主要用来测试 list 的 element.Value.(*类型) 的作用， 其实感觉类似 java 中的类型转换，
// 比如我从 user := element.Value.(*User), 那我的 user 就具备 User 类型的属性和行为
//  user := element.Value.(User) 与上面的区别
// 前者修改是会影响 list 的元素， 后者不会.
func main()  {

	key := "key1"
	var val SS = "value1"
	e := &Entry{key: key, value: val}
	element := l.PushFront(e)
	cache[key] = element

	// (*Entry) 这是干嘛的? reference: https://qa.1r1g.com/sf/ask/1610316081/
	// 如果不加 *， 那么运行:  panic: interface conversion: interface {} is *main.Entry, not main.Entry

	// 相当于把 Element 存储的Value 转成了具体的对象， 也就是 Entry
	// 因为存储的是 Entry 的引用， 因此修改它，也会修改列表中的值。
	// 如果不想联动修改，可以存储 Entry
	fmt.Println("list Element: ", element.Value.(*Entry))
	fmt.Println("map 中 key[", key, "]=", cache[key].Value)
	fmt.Println("list 中： ", l.Front().Value)

	// print result:
	// list==  &{key1 value1}
	// map 中 key[ key1 ]= &{key1 value1}
	// list 中：  &{key1 value1}

}
