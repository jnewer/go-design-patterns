package main

import "fmt"

type Sport interface {
	Run()
}

type Basketball struct {
}
type Football struct {
}

func NewBasketball() *Basketball {
	return &Basketball{}
}

func NewFootbal() *Football {
	return &Football{}
}

func (*Basketball) Run() {
	fmt.Println("play basketball")
}

func (*Football) Run() {
	fmt.Println("play football")
}

type SportFacade struct {
	basketball Basketball
	football   Football
}

func NewSportFacade() *SportFacade {
	return &SportFacade{
		basketball: Basketball{},
		football:   Football{},
	}
}

func (f *SportFacade) PlayBasketball() {
	f.basketball.Run()
}

func (f *SportFacade) PlayFootball() {
	f.football.Run()
}

func main() {
	sf := NewSportFacade()
	sf.PlayBasketball()
	sf.PlayFootball()
}
