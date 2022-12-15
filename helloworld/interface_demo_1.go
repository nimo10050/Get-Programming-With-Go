package main

import "fmt"

type Mover interface {
	Move()
}

type Dog struct {
}

type Cat struct {
}

func (d *Dog) Move() {
	fmt.Println("dog is moving...")
}

func (c Cat) Move() {
	fmt.Println("cat is moving...")
}

func main() {
	var m Mover
	m = Cat{}
	m.Move()

	m = &Dog{}
	m.Move()

}
