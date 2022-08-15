package main

import "fmt"

type Order struct {
	Name string
}

type OrderList struct {
	Orders []Order
}

func NewOrderList() *OrderList {
	orders := make([]Order, 0)
	return &OrderList{
		orders,
	}
}

func (ol *OrderList) GetIterator() func() (Order, bool) {
	index := 0
	return func() (order Order, ok bool) {
		if index >= len(ol.Orders) {
			return
		}

		order, ok = ol.Orders[index], true
		index++
		return
	}
}

func (ol *OrderList) Add(order Order) {
	ol.Orders = append(ol.Orders, order)
}

func main() {
	ol := NewOrderList()
	o1 := Order{"book"}
	o2 := Order{"clothes"}
	ol.Add(o1)
	ol.Add(o2)

	it := ol.GetIterator()
	for {
		order, ok := it()
		if !ok {
			break
		}

		fmt.Println("Order: ", order.Name)
	}
}
