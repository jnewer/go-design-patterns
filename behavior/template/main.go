package main

import "fmt"

type IPerson interface {
}

type Person struct {
}

func (p *Person) Birth() {
	fmt.Println("birth...")
}
func (p *Person) Live() {
}

func (p *Person) Dead() {
	fmt.Println("dead...")
}

func (p *Person) LifeTime() {
	p.Birth()
	p.Live()
	p.Dead()
}

type SuccessPerson struct {
	Person
}

type FailPerson struct {
	Person
}

func (p *SuccessPerson) Live() {
	fmt.Println("success live")
}

func (p *SuccessPerson) LifeTime() {
	p.Birth()
	p.Live()
	p.Dead()
}

func (p *FailPerson) Live() {
	fmt.Println("fail live")
}

func (p *FailPerson) LifeTime() {
	p.Birth()
	p.Live()
	p.Dead()
}

func main() {
	p := Person{}
	sp := SuccessPerson{p}
	fp := FailPerson{p}

	sp.LifeTime()

	fp.LifeTime()
}
