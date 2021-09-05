package main

import "fmt"

type Goods struct {
	ID        uint
	Price     uint
	Name      string
	GoodsType string
}
type Course struct {
	ID        uint
	Price     uint
	Name      string
	GoodsType string
}

type Discount struct {
	Discount uint
	Price    uint
}

type UserOrder struct {
	Id      uint
	Price   uint
	Name    uint
	OrderNo string
	GoodsID uint
}
type Order interface {
	Pay() (uint, error)
	GetPrice(discount uint) (uint, error)
	ContinuePay(orderID uint) error
}

func main() {
	g := Goods{
		ID:        1,
		Price:     10,
		Name:      "apple",
		GoodsType: "fruit",
	}

	c := Course{
		ID:        2,
		Price:     20,
		Name:      "正式课",
		GoodsType: "course",
	}

	s := []Order{&g, &c} //通过指针实现
	for _, o := range s {
		id, _ := o.Pay()
		o.GetPrice(80)
		o.ContinuePay(id)
	}
	//c.Pay()
	//c.GetPrice(80)
	//c.ContinuePay(1)
}

func (g *Goods) Pay() (uint, error) {
	fmt.Printf("%s成功支付了%d元\n", g.Name, g.Price)

	return 1, nil
}

func (g *Goods) GetPrice(discount uint) (uint, error) {
	g.Price = (g.Price * discount) / 100
	return g.Price, nil
}

func (g *Goods) ContinuePay(orderID uint) error {
	fmt.Printf("订单id为：%d,%s成功继续支付了%d元\n", orderID, g.Name, g.Price)
	return nil
}

func (g *Course) Pay() (uint, error) {
	fmt.Printf("%s成功支付了%d元\n", g.Name, g.Price)

	return 1, nil
}

func (g *Course) GetPrice(discount uint) (uint, error) {
	g.Price = (g.Price * discount) / 100
	return g.Price, nil
}

func (g *Course) ContinuePay(orderID uint) error {
	fmt.Printf("订单id为：%d,%s成功继续支付了%d元\n", orderID, g.Name, g.Price)
	return nil
}
