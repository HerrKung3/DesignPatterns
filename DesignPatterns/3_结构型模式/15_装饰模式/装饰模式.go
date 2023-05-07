package main

import "fmt"

//动态地给一个对象添加一些额外的职责。就增加功能来说，此模式比生成子类更加灵活
//例子：通过装饰器给手机加入贴膜和手机壳功能

// === 抽象层 ===

type Phone interface {
	Show()
}

// Decorator 抽象的装饰器，装饰器的基础类（该类本应该是interface，但golang中的interface不允许存在成员属性）
type Decorator struct {
	phone Phone
}

func (d *Decorator) Show() {

}

// === 实现层 ===

type Huawei struct {
}

func (hw *Huawei) Show() {
	fmt.Println("It's a Huawei phone")
}

type Xiaomi struct {
}

func (xm *Xiaomi) Show() {
	fmt.Println("It's a Xiaomi phone")
}

// MoDecorator 具体的装饰器
type MoDecorator struct {
	Decorator
}

func (md *MoDecorator) Show() {
	md.phone.Show()
	//额外功能
	fmt.Println("贴膜中...")
}

func NewMoDecorator(phone Phone) Phone {
	return &MoDecorator{Decorator{
		phone: phone,
	}}
}

// KeDecorator 具体的装饰器
type KeDecorator struct {
	Decorator
}

func (kd *KeDecorator) Show() {
	kd.phone.Show()
	//额外功能
	fmt.Println("装壳中...")
}

func NewKeDecorator(phone Phone) Phone {
	return &KeDecorator{Decorator{
		phone: phone,
	}}
}

// === 逻辑层 ===
func main() {
	var huawei Phone
	huawei = new(Huawei)

	var moHuawei Phone
	moHuawei = NewMoDecorator(huawei)
	moHuawei.Show()

	var KeMoHuawei Phone
	KeMoHuawei = NewKeDecorator(moHuawei)
	KeMoHuawei.Show()

	var xiaomi Phone
	xiaomi = new(Xiaomi)
	var keXiaomi Phone
	keXiaomi = NewKeDecorator(xiaomi)
	keXiaomi.Show()
}
