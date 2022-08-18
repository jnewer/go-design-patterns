package main

import "fmt"

type IDiscount interface {
	Discount() float32
}

type Book struct {
	Price float32
}

func (b *Book) GetPrice(d IDiscount) float32 {
	return b.Price * d.Discount()
}

type Discount85 struct{}

func (d *Discount85) Discount() float32 {
	return 0.85
}

type Discount65 struct{}

func (d Discount65) Discount() float32 {
	return 0.65
}

func main() {
	b := Book{100}
	d65 := Discount65{}
	p := b.GetPrice(&d65)
	fmt.Printf("p: %v\n", p)
	d85 := Discount85{}
	p = b.GetPrice(&d85)
	fmt.Printf("p: %v\n", p)
}
