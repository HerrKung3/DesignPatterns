package main

import "fmt"

// Banker 错误示范
type Banker struct {
}

func (b *Banker) Save() {
	fmt.Println("Save")
}

func (b *Banker) Trans() {
	fmt.Println("Trans")
}

func (b *Banker) Pay() {
	fmt.Println("Pay")
}

// Shares 新增一个股票功能, 给一个类新增功能时，新增的方法可能与其他方法耦合，进而影响到其他方法
func (b *Banker) Shares() {
	fmt.Println("shares")
}
func main() {
	//错误示范
	banker := Banker{}
	banker.Trans()
	banker.Save()
	banker.Pay()
	banker.Shares()
}
