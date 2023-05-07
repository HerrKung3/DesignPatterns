package main

import "fmt"

// GB 适配目标
type GB interface {
	UseGB()
}

// ES 被适配对象
type ES struct {
}

func (es *ES) UseES() {
	fmt.Println("Use European standard")
}

// 适配器 包含被适配对象属性，并实现适配目标接口的方法
type adapter struct {
	es *ES
}

// UseGB 调用接口时实际调用了被适配对象方法(适配器类加了一层实现适配目标的接口方法的马甲，本质是方法重写)
func (a *adapter) UseGB() {
	fmt.Println("Use Chinese standard")
	a.es.UseES()
}

func newAdapter() adapter {
	return adapter{
		new(ES),
	}
}

func main() {
	a := newAdapter()
	a.UseGB()
}
