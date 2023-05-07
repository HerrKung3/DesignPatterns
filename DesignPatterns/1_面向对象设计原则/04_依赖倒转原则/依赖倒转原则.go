package main

import "fmt"

//依赖于抽象(接口)，不要依赖于具体的实现(类)，也就是针对接口编程

//司机张三开奔驰				司机李四开宝马				业务逻辑层
//
//-------------------------------------------------------- 向下依赖抽象层
//
//抽象汽车					抽象司机					抽象层
//抽象启动run接口				抽象驾驶drive接口
//
//-------------------------------------------------------- 向上依赖抽象层
//
//实现run方法	实现run方法	实现驾驶方法	实现驾驶方法	实现层
//奔驰  			宝马 		张三 		李四

// ===== 抽象层 =====

type Car interface {
	Run()
}

type Driver interface {
	Drive(car Car)
}

// ===== 实现层 ======

type BENZ struct {
}

func (b *BENZ) Run() {
	fmt.Println("Benz is running")
}

type BMW struct {
}

func (b *BMW) Run() {
	fmt.Println("BMW is running")
}

type L4 struct {
}

// Drive 向上依赖抽象层
func (l4 *L4) Drive(car Car) {
	fmt.Println("Li3 drive car")
	car.Run()
}

type Z3 struct {
}

func (z3 *Z3) Drive(car Car) {
	fmt.Println("Zhang3 drive car")
	car.Run()
}

// ===== 业务逻辑层 =====
func main() {
	//向下依赖抽象层
	var benz Car
	benz = new(BENZ)
	var l4 Driver
	var bmw Car
	bmw = new(BMW)
	var z3 Driver
	z3 = new(Z3)
	l4 = new(L4)

	l4.Drive(benz)
	l4.Drive(bmw)
	z3.Drive(bmw)
	z3.Drive(benz)
}
