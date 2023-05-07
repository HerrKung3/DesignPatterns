package main

import "fmt"

// Fruit fruit类过于庞大，含有多个判断语句，不利于测试和维护, 违背单一职责原则
type Fruit struct {
}

func (f *Fruit) Show(name string) {
	fmt.Printf("It's a %s\n", name)
}

// NewFruit 新增水果时，需要修改构造函数，违背了开闭原则
func NewFruit(name string) *Fruit {
	fruit := &Fruit{}
	if name == "apple" {
		//do something about apple
	} else {
		//
	}
	return fruit
}

// 导致业务逻辑层依赖基础类模块
func main() {
	apple := NewFruit("apple")
	apple.Show("apple")
	pear := NewFruit("pear")
	pear.Show("pear")
}
