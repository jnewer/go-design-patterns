package main

import (
	"fmt"
)

type Subject interface {
	Do()
}

type RelaSubject struct {
}

func (rs *RelaSubject) Do() {
	fmt.Println("do something...")
}

type Proxy struct {
	RelaSubject
}

func (p *Proxy) Do() {
	fmt.Println("before proxy")
	p.RelaSubject.Do()
	fmt.Println("after proxy")
}

func main() {
	real := RelaSubject{}
	proxy := Proxy{real}
	proxy.Do()
}
