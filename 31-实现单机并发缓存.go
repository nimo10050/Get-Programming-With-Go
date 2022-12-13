package main

import (
	"GeeCache/lru"
	"fmt"
	"sync"
)

// ByteView 紧接上文:
// 为了实现容量大小可控，所以我们存入的每一个元素都要计算大小
// 为了计算大小，存入缓存的 value 必须实现 Length()
// 但是我们为每个 value 类型都实现一个 Length() 方法难免有些不现实
// 所以我们将存入的每个 value 都转化为字节数组

// 考虑到接下来可能要基于存入的 value (字节数组)做一些操作， 我们为了优雅一些。
// 定义一个对象 用来做如下事情： 封装字节数组以及操作字节数组的一些操作

type ByteView struct {
	b []byte
}

// Len 暂时能想到要用到这一个方法
func (v ByteView) Len() int64 {
	return int64(len(v.b))
}

// 在上一章节， 我们设计的缓存系统存在一个明显的问题：
// 当多个线程并发访问读写时， 会出现所所谓的线程安全问题。（如果有 Java 开发经验的话， 可以参考 java.util.HashMap ）
// 比如一个电商系统， 我们可以把库存和商品的数据都放在同一个 cache 对象中。
// 当多个线程并发读写 cache 对象时，为了保证线程安全，我们给 cache 对象加了一把锁。
//
// 因为有个线程需要把库存数据写入缓存，我们给 cache 对象加了锁， 而此时正好有多个线程需要读取商品信息，因此他们需要等待锁释放了， 才能读取到商品信息。
// 在锁释放之前， 这些线程都要排队等待。

// 为了提高缓存系统的并发读写的性能， 我们同样可以参考 Java 中 java.util.concurrent.ConcurrentHashMap 的分段锁的设计思想。
// 将缓存的数据进行分组，库存数据缓存到 store 分组, 商品数据缓存到 item 分组， 如下所示：

// TODO 差一张图

// 这样之后， 当我们一个线程访问库存时， 就只需要 groupName=store 缓存数据给锁起来, 而 groupName=item 的数据不用锁。

// 综上所述， 我们需要朝着如下两个大方向走:
// 1. 搞清楚 go 语言中如何实现加锁，释放锁
// 2. 既然把缓存数据分组了，那么就需要搞一个新的结构， 来定义缓存对象。

// 因为我们每个缓存对象都要持有一把锁 sync.Mutex
// 所以我们不能直接用第一天定义的 lru.Cache, 需要重新定义一个缓存对象， 我们这里命名为 cache, 这里先不要在乎大小写的问题（是否被公开访问）
// 每个缓存对象都需要统计占用的字节大小， 所以我们需要一个 cacheBytes 字段来存储这个值
type cache struct {
	lru *lru.Cache
	mu  sync.Mutex

	cacheBytes int64
}

// 紧接着我们要给这个新的缓存对象， 赋予 CRUD 的能力，比如查询缓存
// 因为锁要公用， 所以这里的 cache 是指针类型
func (c *cache) get(key string) ByteView {
	c.mu.Lock()
	defer c.mu.Unlock()
	if v, ok := c.lru.Get(key); ok {
		return v.(ByteView), ok
	}
	return
}

// Group 紧接着我们定义一个 分组 类型
type Group struct {
	name      string // 分组名称
	mainCache cache  // 单个缓存对象
}

// New 接下来就是定义或者 Group 对象的方法， 显然还是一个 New 方法
// 我们暂时不要关注实现细节， 这个方法暂时这样丢在这里
func NewGroup(name string) (g *Group) {
	return &Group{name: name, mainCache: nil}
}

// 接下来我们需要通过 Group 对象来获取缓存数据
func (g *Group) Get(key string) (ByteView, error) {
	if v, ok := g.mainCache.get(key); ok {
		return v, nil
	}
	return nil, fmt.Errorf("dont get cache.")
}

// 这里我们还漏掉了一个功能，如果缓存未命中， 我们还要提供回调方法。
// 这个回调方法我们可以直接定义在上面的 Get 方法的入参中。也可以放在 Group 对象中。、
// 为了方便， 我们定义在 Group 对象中。

type Group struct {
	name      string // 分组名称
	mainCache cache  // 单个缓存对象
	getter    Getter
}

// Getter 定义函数式接口
type Getter interface {
	Get(key string)
}

type GetterFunc func(key string) ([]byte, error)

func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

// Get 重新定义 Group 的 Get 方法
func (g *Group) Get(key string) (ByteView, error) {
	if v, ok := g.mainCache.get(key); ok {
		return v, nil
	}
	// 从回调方法中取
	return g.getter.Get(key)
}

// 到此为止， 我们已经定义了一个 Group 对象， 用来缓存一组数据，同时也定义了 Group 对象的 Get 方法。
// 那么问题来了， 我们肯定有多个 Group 对象的哇，所以我们就需要一个 Group 数组来存放这多个 Group 对象

var groups = make(map[string]*Group)

// 那么问题又来了，既然有多个 Group 对象， 那肯定又涉及到并发读写这些对象。这时候我们就要考虑用读写锁来解决这个问题了。
var rwMutex = sync.RWMutex

// 紧接着定义读写 group 的方法， 比如 NewGroup：
// 我们对上面的定义的 NewGroup 方法进行补充:
func MewGroup(name string, cacheBytes int64) *Group {
	rwMutex.Lock()
	defer rwMutex.Unlock()
	g := &Group{name: name, mainCache: cache{cacheBytes: cacheBytes}}
	groups[name] = g
	return g
}

// 获取 Group 对象的方法
func GetGroup(name string) *Group {
	rwMutex.Lock()
	defer rwMutex.Unlock()
	g := groups[name]
	return g
}
