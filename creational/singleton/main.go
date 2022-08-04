package main

import (
	"fmt"
	"sync"
)

type Singleton interface {
	dosomething()
}
type singleton struct {
}

func (s *singleton) dosomething() {
	fmt.Println("do some thing")
}

var (
	once     sync.Once
	instance *singleton
)

func getInstance() Singleton {
	once.Do(
		func() {
			instance = &singleton{}
		},
	)
	return instance
}
func main() {
	s1 := getInstance()
	fmt.Printf("s1: %p\n", s1)
	s2 := getInstance()
	fmt.Printf("s2: %p\n", s2)
}
