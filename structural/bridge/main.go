package main

import (
	"fmt"
)

type DrawAPI interface {
	DrawCircle(radius, x, y int)
}
type RedCircle struct{}

func NewRedCircle() *RedCircle {
	return &RedCircle{}
}
func (rc *RedCircle) DrawCircle(radius, x, y int) {
	fmt.Printf("Drawing Circle[ color: red, radius: %d, x: %d, y: %d ]\n", radius, x, y)
}

type GreenCircle struct{}

func NewGreenCircle() *GreenCircle {
	return &GreenCircle{}
}
func (gc *GreenCircle) DrawCircle(radius, x, y int) {
	fmt.Printf("Drawing Circle[ color: green, radius: %d, x: %d, y: %d ]\n", radius, x, y)
}

type ShapeCircle struct {
	Radius  int
	X       int
	Y       int
	drawAPI DrawAPI
}

func NewShapeCircle(radius, x, y int, drawAPI DrawAPI) *ShapeCircle {
	return &ShapeCircle{
		Radius:  radius,
		X:       x,
		Y:       y,
		drawAPI: drawAPI,
	}
}

func (sc *ShapeCircle) Draw() {
	sc.drawAPI.DrawCircle(sc.Radius, sc.X, sc.Y)
}

func main() {
	redCircle := NewShapeCircle(5, 6, 8, NewRedCircle())
	redCircle.Draw() //Drawing Circle[ color: red, radius: 5, x: 6, y: 8 ]

	greenCircle := NewShapeCircle(1, 2, 4, NewGreenCircle())
	greenCircle.Draw() //Drawing Circle[ color: green, radius: 1, x: 2, y: 4 ]
}
