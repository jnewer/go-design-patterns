package main

import "fmt"

type OldInterface interface {
	OldMethod()
}
type OldImpl struct {
}

func (OldImpl) OldMethod() {
	fmt.Println("OldMethod")
}

type NewInterface interface {
	NewMethod()
}

type Adapter struct {
	OldInterface
}

func (a *Adapter) NewMethod() {
	fmt.Println("NewMethod")
	a.OldMethod()
}

func main() {
	oldInterface:=OldImpl{}
	a:=Adapter{OldInterface: oldInterface}
	a.NewMethod()
}
