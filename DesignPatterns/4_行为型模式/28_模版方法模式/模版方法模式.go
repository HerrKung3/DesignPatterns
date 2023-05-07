package main

import "fmt"

//得子类可以不改变一个算法的结构即可重定义该算法的某些特定步骤

//冲泡咖啡				冲泡茶叶				抽象
//1. 煮水 				1. 煮水 				1. 煮水
//2. 冲泡咖啡			2. 冲泡茶叶			2. 冲泡
//3. 倒入杯中			3. 倒入杯中			3. 倒入杯中
//4. 加糖和牛奶			4. 加柠檬			4. 加料

// Beverage 抽象类， 制作饮料，包裹一个模版，全部实现的步骤
type Beverage interface {
	BoilWater()
	Brew()
	PourInCup()
	AddThings()
	// WantAddThings 是否添加佐料的Hook
	WantAddThings() bool
}

// 封装一套流程模版基类，让具体的制作流程继承且实现
type template struct {
	b Beverage
}

func (t *template) MakeBeverage() {
	if t == nil {
		return
	}
	t.b.BoilWater()
	t.b.Brew()
	t.b.PourInCup()
	if t.b.WantAddThings() {
		t.b.AddThings()
	}
}

// MakeCoffee 制作咖啡
type MakeCoffee struct {
	template
}

func NewMakeCoffee() *MakeCoffee {
	makeCoffee := new(MakeCoffee)
	//b是Beverage，是MakeCoffee的接口，这里需要给接口赋值，让b指向具体的子类，来触发b的全部方法的多态特征
	makeCoffee.b = makeCoffee
	return makeCoffee
}

func (mc *MakeCoffee) BoilWater() {
	fmt.Println("Boil water")
}

func (mc *MakeCoffee) Brew() {
	fmt.Println("Brew coffee")
}

func (mc *MakeCoffee) PourInCup() {
	fmt.Println("Pour in cup")
}

func (mc *MakeCoffee) WantAddThings() bool {
	return true
}

func (mc *MakeCoffee) AddThings() {
	fmt.Println("Add milk ans sugar in coffee")
}

// MakeTea 制作茶
type MakeTea struct {
	template
}

func NewMakeTea() *MakeTea {
	makeTea := new(MakeTea)
	makeTea.b = makeTea
	return makeTea
}

func (mt *MakeTea) BoilWater() {
	fmt.Println("Boil water")
}

func (mt *MakeTea) Brew() {
	fmt.Println("Brew tea")
}

func (mt *MakeTea) PourInCup() {
	fmt.Println("Pour in cup")
}

func (mt *MakeTea) AddThings() {
	fmt.Println("Add lemon in tea")
}

func (mt *MakeTea) WantAddThings() bool {
	return false
}

func main() {
	//制作咖啡
	makeCoffee := NewMakeCoffee()
	makeCoffee.MakeBeverage()
	//制作茶
	makeTea := NewMakeTea()
	makeTea.MakeBeverage()
}
