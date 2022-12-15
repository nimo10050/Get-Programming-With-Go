package main

import "fmt"

type Distance struct {
	x, y float64
}

func diff(p, q Distance) float64 {
	return p.x - q.x
}

func (p Distance) compare(q Distance) float64 {
	return p.x - q.x
}

func (p *Distance) compare1(q Distance) float64 {
	return p.x - q.x
}

// SetX 值类型的接收者
func (p Distance) SetX(x float64) {
	p.x = x
}

// SetX2 指针类型的接收者
func (p *Distance) SetX2(x float64) {
	p.x = x
}

func main() {
	//fmt.Println(diff(Distance{2, 1}, Distance{1, 1}))
	p := Distance{3, 4}
	//fmt.Println(p.compare(Distance{1, 1}))
	p.SetX(6)
	fmt.Println(p.x)

	(&p).SetX2(6)
	fmt.Println(p.x)
	// p1 := Distance{2, 6}
	//fmt.Println(p1.compare1(Distance{1, 1}))
	//fmt.Println((&p1).compare1(Distance{1, 1}))

}
