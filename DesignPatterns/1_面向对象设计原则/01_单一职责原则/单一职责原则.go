package main

import "fmt"

//类的职责单一，对外只提供一种功能，而引起变化原因都应该只有一个

// Clothes 错误示范
type Clothes struct {
}

func (c *Clothes) OnWorking() {
	fmt.Println("Clothes style for work")
}

func (c *Clothes) OnShopping() {
	fmt.Println("Clothes style for shopping")
}

// ClothesForShop 正确示范
type ClothesForShop struct {
}
type ClothesForWork struct {
}

func (cs *ClothesForShop) Style() {
	fmt.Println("Clothes style for shopping")
}

func (cs *ClothesForWork) Style() {
	fmt.Println("Clothes style for work")
}

func main() {
	//错误示范
	c := Clothes{}
	c.OnWorking()
	c.OnShopping()
	//正确示范
	cw := ClothesForWork{}
	cw.Style()
	cs := ClothesForShop{}
	cs.Style()
}
