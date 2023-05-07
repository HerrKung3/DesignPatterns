package main

import "fmt"

////通过专门定义一个类来负责创建其他的类的实例，被创建的类通常都具有共同的父类(抽象层)

//优点：实现了对象的创建与使用分离
//缺点：工厂职责过重，一旦无法正常工作，系统会受影响
//     违反开闭原则，添加新类时，需要修改工厂类的代码
//适用场景：工厂类负责创建的对象较少时
//		  客户端只知道传入工厂的参数，对如何创建对象并不关心

// === 抽象层 ===

type Fruits interface {
	Show()
}

// === 实现层 ===

type Apple struct {
}

func (a *Apple) Show() {
	fmt.Println("It's a apple")
}

type Pear struct {
}

func (p *Pear) Show() {
	fmt.Println("Its a pear")
}

// Factory Mode
type Factory struct {
}

func (f *Factory) CreateFruit(name string) Fruits {
	//符合多态的赋值，父类指针指向子类对象
	var fruits Fruits
	if name == "apple" {
		fruits = new(Apple)
	} else if name == "pear" {
		fruits = new(Pear)
	}
	return fruits
}

// 业务逻辑层
func main() {
	factory := new(Factory)
	apple := factory.CreateFruit("apple")
	apple.Show()
	pear := factory.CreateFruit("pear")
	pear.Show()
}
