package main

import "fmt"

type sol int

type report struct {
	sol
}

func (s sol) days(s1 int) {
	diff := int(s) - s1
	fmt.Println(diff)
}

func main() {
	r := report{1}
	r.days(0)
}
