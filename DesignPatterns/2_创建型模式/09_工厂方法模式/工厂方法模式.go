package main

import "fmt"

//简单工厂模式+开闭原则=工厂模式

// === 抽象层 ===

type Fruit interface {
	Show()
}

type Factory interface {
	Create() Fruit
}

// === 实现层 ===

type Apple struct {
}

type Pear struct {
}

type AppleFactory struct {
}

type PearFactory struct {
}

func (a *Apple) Show() {
	fmt.Println("It's a apple")
}

func (p *Pear) Show() {
	fmt.Println("It's a pear")
}

func (af *AppleFactory) Create() Fruit {
	var fruit Fruit
	fruit = new(Apple)
	return fruit
}

func (pf *PearFactory) Create() Fruit {
	var fruit Fruit
	fruit = new(Pear)
	return fruit
}

// 业务逻辑层
func main() {
	var appleFactory Factory
	appleFactory = new(AppleFactory)
	apple := appleFactory.Create()
	apple.Show()

	var pearFactory Factory
	pearFactory = new(PearFactory)
	pear := pearFactory.Create()
	pear.Show()
}
