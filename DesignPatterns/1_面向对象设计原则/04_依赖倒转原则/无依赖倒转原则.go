package main

import "fmt"

//错误示范，耦合度极高，代码重复度高, 难以维护
//drivers:ZhangSan LiSi
//cars: benz bmw

type Benz struct {
}

type Bmw struct {
}

type Zhang3 struct {
}

type Li4 struct {
}

func (b *Benz) Run() {
	fmt.Println("Benz is running")
}

func (b *Bmw) Run() {
	fmt.Println("BMW is running")
}

func (z3 *Zhang3) DriveBenz(benz *Benz) {
	fmt.Println("Zhang3 is driving Benz")
	benz.Run()
}

func (z3 *Zhang3) DriveBmw(bmw *Bmw) {
	fmt.Println("Zhang3 is driving BMW")
	bmw.Run()
}

func (l4 *Li4) DriveBenz(benz *Benz) {
	fmt.Println("Li4 is driving Benz")
	benz.Run()
}

func (l4 *Li4) DriveBmw(bmw *Bmw) {
	fmt.Println("Li4 is driving BMW")
	bmw.Run()
}

func main() {
	benz := &Benz{}
	bmw := &Bmw{}
	z3 := Zhang3{}
	l4 := Li4{}
	z3.DriveBenz(benz)
	z3.DriveBmw(bmw)
	l4.DriveBenz(benz)
	l4.DriveBmw(bmw)
}
