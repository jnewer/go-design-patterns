package main

import (
	"fmt"
)

type Shape interface {
	Draw()
}
type Circle struct {
	X      int
	Y      int
	Radius int
	Color  string
}

func NewCircle(color string) *Circle {
	return &Circle{
		Color: color,
	}
}

func (c *Circle) SetX(x int) {
	c.X = x
}

func (c *Circle) SetY(y int) {
	c.Y = y
}

func (c *Circle) SetRadius(radius int) {
	c.Radius = radius
}

func (c *Circle) Draw() {
	fmt.Printf("draw [color: %s, x: %d, y: %d, radius: %d] \n", c.Color, c.X, c.Y, c.Radius)
}

type ShapeFactory struct {
	circleMap map[string]Shape
}

func NewShapeFactory() *ShapeFactory {
	return &ShapeFactory{
		circleMap: make(map[string]Shape),
	}
}

func (sf *ShapeFactory) GetCircle(color string) Shape {
	circle := sf.circleMap[color]

	if circle == nil {
		circle = NewCircle(color)
		sf.circleMap[color] = circle
		fmt.Println("circle color: ", color)
	}

	return circle
}

func main() {
	sf := NewShapeFactory()
	sf.GetCircle("red").Draw()
	sf.GetCircle("red").Draw()
	sf.GetCircle("green").Draw()
	sf.GetCircle("green").Draw()
	sf.GetCircle("blue").Draw()
	sf.GetCircle("blue").Draw()
}
