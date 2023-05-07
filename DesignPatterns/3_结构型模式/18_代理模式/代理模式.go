package main

import "fmt"

//为其他对象提供一种代理已控制对这个对象的访问

type Goods struct {
	Kind string //种类
	Fact bool   //真伪
}

// === 抽象层 ===

type Shopping interface {
	Buy(goods *Goods)
}

// === 实现层 ===

type KoreaShopping struct {
}

func (ks *KoreaShopping) Buy(good *Goods) {
	fmt.Printf("Buy %s in Korea\n", good.Kind)
}

type USAShopping struct {
}

func (us *USAShopping) Buy(good *Goods) {
	fmt.Printf("Buy %s in USA\n", good.Kind)
}

// OverseaProxy oversea proxy
type OverseaProxy struct {
	shopping Shopping
}

func NewOverseaProxy(shopping Shopping) Shopping {
	return &OverseaProxy{shopping}
}

// Buy 增强buy方法
func (op *OverseaProxy) Buy(goods *Goods) {
	//1 +辨别真伪
	if op.distinguish(goods) {
		//2 buy
		op.shopping.Buy(goods)
		//3 +海关安检
		op.check(goods)
	} else {
		fmt.Printf("%s is a fake", goods.Kind)
	}
}

func (op *OverseaProxy) distinguish(goods *Goods) bool {
	return goods.Fact
}

func (op *OverseaProxy) check(goods *Goods) {
	fmt.Printf("checking %s\n", goods.Kind)
}

// === 业务逻辑层 ===
func main() {
	g1 := Goods{
		Kind: "facial mask",
		Fact: true,
	}
	g2 := Goods{
		Kind: "CET-4",
		Fact: false,
	}

	var KShopping Shopping
	KShopping = new(KoreaShopping)
	var KProxy Shopping
	KProxy = NewOverseaProxy(KShopping)
	KProxy.Buy(&g1)

	var UShopping Shopping
	UShopping = new(USAShopping)
	var UProxy Shopping
	UProxy = NewOverseaProxy(UShopping)
	UProxy.Buy(&g2)
}
