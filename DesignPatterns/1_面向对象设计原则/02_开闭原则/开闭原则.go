package main

import "fmt"

//类的改动是通过增加代码来进行的，而不是修改源代码
//定义一个抽象层接口，其中实现抽象的方法
//实现层定义类时，类实现抽象层接口的方法并重写
//当需要新增业务时，通过定义一个实现接口层方法的类来达到目的，而不是像平铺设计那样，直接在原有类中新增方法

// AbstractBanker 抽象的银行业务员
type AbstractBanker interface {
	DoBus()
}

type SaveBanker struct {
}

type TransBanker struct {
}

type PayBanker struct {
}

// SharesBanker 新增一个股票功能
type SharesBanker struct {
}

func (sb SaveBanker) DoBus() {
	fmt.Println("Save")
}

func (tb TransBanker) DoBus() {
	fmt.Println("Trans")
}

func (pb PayBanker) DoBus() {
	fmt.Println("Pay")
}

func (shb *SharesBanker) DoBus() {
	fmt.Println("Shares")
}

// BankBus 实现一个架构层, 针对interface接口进行封装
func BankBus(banker AbstractBanker) {
	//通过接口向下调用(多态现象)
	banker.DoBus()
}

func main() {
	BankBus(&SaveBanker{})
	BankBus(&PayBanker{})
	BankBus(&TransBanker{})
	BankBus(&SharesBanker{})
}
