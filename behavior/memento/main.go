package main

import "fmt"

type Memento struct {
	state string
}

func NewMemento(state string) *Memento {
	return &Memento{state: state}
}

func (m *Memento) GetState() string {
	return m.state
}

func (m *Memento) SetState(state string) {
	m.state = state
}

type Originator struct {
	state string
}

func NewOriginator(state string) *Originator {
	return &Originator{state: state}
}

func (o *Originator) GetState() string {
	return o.state
}

func (o *Originator) SetState(state string) {
	o.state = state
}

func (o *Originator) SaveStateToMemento() *Memento {
	return &Memento{o.state}
}

func (o *Originator) GetStateFromMemento(memento *Memento) {
	o.state = memento.GetState()
}

type CareTaker struct {
	MementoList map[int]*Memento
}

func NewCareTaker() *CareTaker {
	return &CareTaker{MementoList: make(map[int]*Memento)}
}

func (ct *CareTaker) Add(index int, memento *Memento) {
	ct.MementoList[index] = memento
}

func (ct *CareTaker) Get(index int) *Memento {
	return ct.MementoList[index]
}

func main() {
	careTaker := NewCareTaker()
	originator := NewOriginator("state 1")
	originator.SetState("state 2")
	careTaker.Add(1, originator.SaveStateToMemento())
	originator.SetState("state 3")
	careTaker.Add(2, originator.SaveStateToMemento())
	originator.SetState("state 4")

	fmt.Println("current state:", originator.GetState())
	originator.GetStateFromMemento(careTaker.Get(1))
	fmt.Println("first state:", originator.GetState())
	originator.GetStateFromMemento(careTaker.Get(2))
	fmt.Println("second state:", originator.GetState())
}
