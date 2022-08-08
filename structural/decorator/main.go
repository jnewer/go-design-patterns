package main

import "fmt"

type Shape interface {
	Draw()
}

type Circle struct{}

func NewCircle() *Circle {
	return &Circle{}
}

func (c *Circle) Draw() {
	fmt.Println("Draw circle")
}

type RedShapeDecorator struct {
	DecoratedShap Shape
}

func NewRedShapeDecorator(decShape Shape) *RedShapeDecorator {
	return &RedShapeDecorator{
		DecoratedShap: decShape,
	}
}

func (rs *RedShapeDecorator) SetRedBorder() {
	fmt.Println("red border")
}

func (rs *RedShapeDecorator) Draw() {
	rs.DecoratedShap.Draw()
	rs.SetRedBorder()
}

func main() {
	c := NewCircle()
	rsd := NewRedShapeDecorator(c)
	rsd.Draw()
}
