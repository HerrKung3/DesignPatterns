package main

import "fmt"

//将一个类的接口转换成客户希望的另外一个接口。使得原本由于接口不兼容而不能一起工作的那些类可以一起工作
//例子：通过适配器将电压从220V转化成5V

// V5 适配目标
type V5 interface {
	Use5V()
}

// V220 被适配
type V220 struct {
}

func (v *V220) Use220V() {
	fmt.Println("Use 220V voltage")
}

// Adapter 适配器类
type Adapter struct {
	v220 *V220
}

func NewAdapter(v220 *V220) *Adapter {
	return &Adapter{
		v220: v220,
	}
}

// Use5V 调用接口时实际调用了被适配对象方法(适配器类加了一层实现适配目标的接口方法的马甲，本质是方法重写)
func (a *Adapter) Use5V() {
	fmt.Println("Charging with 5V")
	a.v220.Use220V()
}

// Phone 业务类
type Phone struct {
	v V5
}

func NewPhone(v V5) *Phone {
	return &Phone{
		v: v,
	}
}

func (p *Phone) Charge() {
	fmt.Println("Phone is charging")
	p.v.Use5V()
}

func main() {
	phone := NewPhone(NewAdapter(new(V220)))
	phone.Charge()
}
