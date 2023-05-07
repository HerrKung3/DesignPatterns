package main

import "fmt"

type SaleStrategy interface {
	Sale(float64)
}

type TwentyPercentDiscount struct {
}

func (t *TwentyPercentDiscount) Sale(n float64) {
	fmt.Println("20% discount: ", n*0.8)
}

type MoneyOff struct {
}

func (m *MoneyOff) Sale(n float64) {
	if n >= 200 {
		fmt.Println("Subtract 100 from 200: ", int(n)-(int(n)/200)*100)
	}
}

// Custom 环境类 发起策略调用
type Custom struct {
	saleStrategy SaleStrategy
}

func (c *Custom) SetSaleStrategy(s SaleStrategy) {
	c.saleStrategy = s
}

func (c *Custom) UseSaleStrategy(n float64) {
	c.saleStrategy.Sale(n)
}

func main() {
	c := new(Custom)
	c.SetSaleStrategy(new(MoneyOff))
	c.UseSaleStrategy(300)
	c.SetSaleStrategy(new(TwentyPercentDiscount))
	c.UseSaleStrategy(300)
}
