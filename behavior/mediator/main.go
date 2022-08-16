package main

import (
	"fmt"
	"time"
)

type UnitedNations struct{}

func NewUnitedNations() *UnitedNations {
	return &UnitedNations{}
}

var un = NewUnitedNations()

type Nation struct {
	Name string
}

func NewNation(name string) *Nation {
	return &Nation{name}
}

func (un *UnitedNations) ShowMessage(nation *Nation, msg string) {
	fmt.Printf("%s: [ %s ]: %s \n",
		time.Now().Format("2006-01-02 15:04:05"),
		nation.Name,
		msg)
}

func (n *Nation) ShowMessage(msg string) {
	un.ShowMessage(n, msg)
}

func main() {
	n1 := NewNation("中国")
	n1.ShowMessage("强大")

	n2 := NewNation("美国")
	n2.ShowMessage("嚣张")
}
