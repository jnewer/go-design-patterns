package main

import "fmt"

type Shape interface {
	Draw()
}

type Color interface {
	Fill()
}

type Circle struct{}
type Square struct{}

func (c Circle) Draw() {
	fmt.Println("draw circle")
}

func (s Square) Draw() {
	fmt.Println("draw square")
}

type Red struct{}
type Green struct{}

func (r Red) Fill() {
	fmt.Println("fill red")
}

func (g Green) Fill() {
	fmt.Println("fill green")
}

type AbstractFactory interface {
	GetShape(shapeName string) Shape
	GetColor(colorName string) Color
}

type ShapeFactory struct{}

type ColorFactory struct{}

func (sf ShapeFactory) GetShape(shapeName string) Shape {
	switch shapeName {
	case "circle":
		return &Circle{}
	case "square":
		return &Square{}
	default:
		return nil
	}
}

func (sf ShapeFactory) GetColor(colorName string) Color {
	return nil
}

func (cf ColorFactory) GetShape(shapeName string) Shape {
	return nil
}

func (cf ColorFactory) GetColor(colorName string) Color {
	switch colorName {
	case "red":
		return &Red{}
	case "green":
		return &Green{}
	default:
		return nil
	}
}

type FactoryProducer struct{}

func (fp FactoryProducer) GetFactory(factoryName string) AbstractFactory {
	switch factoryName {
	case "color":
		return &ColorFactory{}
	case "shape":
		return &ShapeFactory{}
	default:
		return nil
	}
}

func NewFactoryProducer() *FactoryProducer {
	return &FactoryProducer{}
}

func main() {
	superFactory := NewFactoryProducer()
	colorFactory := superFactory.GetFactory("color")
	shapeFactory := superFactory.GetFactory("shape")

	red := colorFactory.GetColor("red")
	green := colorFactory.GetColor("green")

	circle := shapeFactory.GetShape("circle")
	square := shapeFactory.GetShape("square")

	circle.Draw()
	red.Fill()

	square.Draw()
	green.Fill()
}
