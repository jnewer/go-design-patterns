package main

import "fmt"

type Reader interface {
	Update()
}

type Reader1 struct {
	subject *Subject
}

func (r *Reader1) Reader1Observer(subject *Subject) {
	r.subject = subject
	r.subject.Attach(r)
}

func (r *Reader1) Update() {
	fmt.Println("reader1: ", r.subject.GetState())
}

type Reader2 struct {
	subject *Subject
}

func (r *Reader2) Reader2Observer(subject *Subject) {
	r.subject = subject
	r.subject.Attach(r)
}
func (r *Reader2) Update() {
	fmt.Println("reader2: ", r.subject.GetState())
}

type Reader3 struct {
	subject *Subject
}

func (r *Reader3) Reader1Observer(subject *Subject) {
	r.subject = subject
	r.subject.Attach(r)
}
func (r *Reader3) Update() {
	fmt.Println("reader3: ", r.subject.GetState())
}

type Subject struct {
	readers []Reader
	state   int
}

func NewSubject() *Subject {
	return &Subject{
		state:   0,
		readers: make([]Reader, 0),
	}
}
func (s *Subject) GetState() int {
	return s.state
}

func (s *Subject) SetState(state int) {
	s.state = state
	s.NotifyAllObservers()
}

func (s *Subject) Attach(observer Reader) {
	s.readers = append(s.readers, observer)
}

func (s *Subject) NotifyAllObservers() {
	for _, observer := range s.readers {
		observer.Update()
	}
}
func main() {
	subject := NewSubject()

	reader1 := Reader1{subject: subject}
	reader2 := Reader2{subject: subject}
	reader3 := Reader3{subject: subject}

	subject.Attach(&reader1)
	subject.Attach(&reader2)
	subject.Attach(&reader3)

	subject.SetState(1)
	subject.SetState(2)
}
