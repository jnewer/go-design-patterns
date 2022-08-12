package main

import (
	"container/list"
	"fmt"
)

type Order interface {
	Execute()
}

type Stock struct {
	Name     string
	Quantity int
}

func NewStock() *Stock {
	return &Stock{
		Name:     "茅台",
		Quantity: 1000,
	}
}

func (s *Stock) Buy() {
	fmt.Printf("股票 [ 名称: %s, 数量: %d ] 买.\n",
		s.Name,
		s.Quantity)
}

func (s *Stock) Sell() {
	fmt.Printf("股票 [ 名称: %s, 数量: %d ] 卖.\n",
		s.Name,
		s.Quantity)
}

type BuyStock struct {
	StockToBy Stock
}

func NewBuyStock(st Stock) *BuyStock {
	return &BuyStock{
		StockToBy: st,
	}
}

func (bs *BuyStock) Execute() {
	bs.StockToBy.Buy()
}

type SellStock struct {
	StockToSell Stock
}

func NewSellStock(st Stock) *SellStock {
	return &SellStock{
		StockToSell: st,
	}
}

func (ss *SellStock) Execute() {
	ss.StockToSell.Sell()
}

type Broker struct {
	OrderList *list.List
}

func NewBroker() *Broker {
	return &Broker{
		OrderList: list.New(),
	}
}

func (b *Broker) TakeOrder(order Order) {
	b.OrderList.PushBack(order)
}

func (b *Broker) PlaceOrders() {
	for i := b.OrderList.Front(); i != nil; {
		nextOrder := i.Next()
		i.Value.(Order).Execute()
		b.OrderList.Remove(i)
		i = nextOrder
	}
}

func main() {
	maotaiStock := NewStock()
	buyStockOrder := NewBuyStock(*maotaiStock)
	sellStockOrder := NewSellStock(*maotaiStock)

	broker := NewBroker()
	broker.TakeOrder(buyStockOrder)
	broker.TakeOrder(sellStockOrder)

	broker.PlaceOrders()
}
