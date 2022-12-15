package main

import "sync"

// 缓存击穿：缓存失效了， 请求打到了 DB 上
// 防止缓存击穿

// 表示正在处理的请求
type call struct {
	wg  sync.WaitGroup // 类似 Java 中的 CountdownLatch
	val interface{}
	err error
}

type Group struct {
	m map[string]*call
}

func (g *Group) Do(key string) (val interface{}) {
	// 如果请求正在处理中, 直接返回
	c := g.m[key]
	if c != nil {
		c.wg.Wait()
		return c.val
	}

	// 创建一个正在处理的请求
	c = new(call)
	c.wg.Add(1)
	g.m[key] = c
	// TODO 获取缓存
	c.wg.Done()

	// 请求处理完， 删除掉
	delete(g.m, key)

	return c.val
}
