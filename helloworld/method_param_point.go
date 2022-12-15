package main

import "fmt"

// 值传递 引用传递？？？

type Point struct {
	x, y float64
}

func testParamIsObj(p Point)  {
	p.x ++
	fmt.Println(p.x)
	p.y ++
}

func testParamIsPoint(p *Point) {
	p.x ++
	p.y ++
}

func main() {
	p := Point{1, 1}
	testParamIsObj(p)
	fmt.Println(p.x)
	testParamIsPoint(&p)
	fmt.Println(p.x)
}
