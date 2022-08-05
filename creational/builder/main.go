package main

import (
	"fmt"
)

type Builder interface {
	buildDisk()
	buildCPU()
	buildRom()
}

type SuperComputer struct {
	Name string
}

func (superComputer *SuperComputer) buildDisk() {
	fmt.Println("SuperComputer buildDisk")
}
func (superComputer *SuperComputer) buildCPU() {
	fmt.Println("SuperComputer buildCPU")
}

func (superComputer *SuperComputer) buildRom() {
	fmt.Println("SuperComputer buildRom")
}

type LowComputer struct {
	Name string
}

func (lowComputer *LowComputer) buildDisk() {
	fmt.Println("LowComputer buildDisk")
}
func (lowComputer *LowComputer) buildCPU() {
	fmt.Println("LowComputer buildCPU")
}

func (lowComputer *LowComputer) buildRom() {
	fmt.Println("LowComputer buildRom")
}

type Director struct {
	builder Builder
}

func NewConstruct(b Builder) *Director {
	return &Director{
		builder: b,
	}
}

func (director *Director) Construct() {
	director.builder.buildDisk()
	director.builder.buildCPU()
	director.builder.buildRom()
}

func main() {
	sc := SuperComputer{}
	d := NewConstruct(&sc)
	d.Construct()
	fmt.Println("--------------")

	lc := LowComputer{}
	d2 := NewConstruct(&lc)
	d2.Construct()
}
