package main

import "fmt"

//工厂模式下：虽然符合开闭原则，产品类数量和工厂类数量相同，随着产品类的数量增加，工厂类数量也在增加
//抽象工厂模式是针对产品族进行产品来生产的，一般来说，要求一个产品族中，产品类型较固定

//							一个产品等级结构
//						  	  ^
//						  	  |
//美国工厂	CPU	NIC	Screen	美国产
//日本工厂	CPU	NIC	Screen	日本产
//中国工厂	CPU	NIC	Screen	中国产	——> 一个产品族

//=== 抽象层 ===

type AbstractCPU interface {
	ShowCPU()
}

type AbstractNIC interface {
	ShowNIC()
}

type AbstractScreen interface {
	ShowScreen()
}

type AbstractFactory interface {
	CreateCPU() AbstractCPU
	CreateNIC() AbstractNIC
	CreateScreen() AbstractScreen
}

//=== 实现层 ===
//中国产品族

type ChinaCPU struct {
}

func (cc *ChinaCPU) ShowCPU() {
	fmt.Println("CPU made in China")
}

type ChinaNIC struct {
}

func (cn *ChinaNIC) ShowNIC() {
	fmt.Println("NIC made in China")
}

type ChinaScreen struct {
}

func (cs *ChinaScreen) ShowScreen() {
	fmt.Println("Screen made in China")
}

type ChinaFactory struct {
}

func (cf *ChinaFactory) CreateCPU() AbstractCPU {
	var cpu AbstractCPU
	cpu = new(ChinaCPU)
	return cpu
}

func (cf *ChinaFactory) CreateNIC() AbstractNIC {
	var nic AbstractNIC
	nic = new(ChinaNIC)
	return nic
}

func (cf *ChinaFactory) CreateScreen() AbstractScreen {
	var screen AbstractScreen
	screen = new(ChinaScreen)
	return screen
}

//日本产品族

type JapanCPU struct {
}

func (jc *JapanCPU) ShowCPU() {
	fmt.Println("CPU made in Japan")
}

type JapanNIC struct {
}

func (jn *JapanNIC) ShowNIC() {
	fmt.Println("NIC made in Japan")
}

type JapanScreen struct {
}

func (js *JapanScreen) ShowScreen() {
	fmt.Println("Screen made in Japan")
}

type JapanFactory struct {
}

func (jf *JapanFactory) CreateCPU() AbstractCPU {
	var cpu AbstractCPU
	cpu = new(JapanCPU)
	return cpu
}

func (jf *JapanFactory) CreateNIC() AbstractNIC {
	var nic AbstractNIC
	nic = new(JapanNIC)
	return nic
}

func (jf *JapanFactory) CreateScreen() AbstractScreen {
	var screen AbstractScreen
	screen = new(JapanScreen)
	return screen
}

//美国产品族同理
//...

//=== 逻辑层 ===

func main() {
	//需求一：需要中国的CPU，日本的网卡和显示器
	var cFac AbstractFactory
	cFac = new(ChinaFactory)
	cCPU := cFac.CreateCPU()
	cCPU.ShowCPU()
	var jFac AbstractFactory
	jFac = new(JapanFactory)
	jNIC := jFac.CreateNIC()
	jNIC.ShowNIC()
	jScreen := jFac.CreateScreen()
	jScreen.ShowScreen()
}
