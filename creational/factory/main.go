package main

import "fmt"

type Fruit interface {
	grant()
	pick()
}

type Apple struct {
}

func (*Apple) grant() {
	fmt.Println("grant apple")
}
func (*Apple) pick() {
	fmt.Println("pick apple")
}

type Orange struct {
}

func (*Orange) grant() {
	fmt.Println("grant orange")
}
func (*Orange) pick() {
	fmt.Println("pick orange")
}
func newFruit(t int) Fruit {
	switch t {
	case 1:
		return &Apple{}
	case 2:
		return &Orange{}
	}
	return nil
}
func main() {
	apple := newFruit(1)
	apple.grant()
	apple.pick()
	fmt.Println("-----------")
	orange := newFruit(2)
	orange.grant()
	orange.pick()
}
