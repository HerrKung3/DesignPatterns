package main

import "fmt"

//定义对象间一种一对多的依赖关系，当一个对象的状态发生改变时，所有依赖于它的对象都得到通知并自动更新

//观察者模式优点：
//			 1 可以实现表示层和数据逻辑层的分离 定义了稳定的消息更新传递机制，并抽象了更新接口
//			 2 在观察者和观察目标之间建立了抽象的耦合，观察目标(被观察者)只需要维持一个抽象观察者集合，无需了解具体观察者
//			 3 支持广播通信，简化了一对多的系统设计难度
//			 4 满足开闭原则
//观察者模式缺点：
//			 1 如果观察者过多，会导致通知比较花时间
//			 2 如果观察者与被观察者之间存在循环依赖，则导致系统崩溃
//			 3 没有相应机制让观察者知道被观察者发生了什么变化，只知道发生了变化

//例子：学生捣乱，班长放哨，学生是观察者，班长是被观察者, 也是观察老师的观察者

// === 抽象层 ===

// Listener 抽象观察者
type Listener interface {
	OnTeacherComing()
}

// Notifier 抽象被观察者
type Notifier interface {
	AddListener(listener Listener)
	RemoveListener(listener Listener)
	Notify()
}

// === 实现层 ===

type Stu1 struct {
	BadThing string
}

func (s1 *Stu1) OnTeacherComing() {
	fmt.Printf("stu1 Stoping doing %s\n", s1.BadThing)
}

type Stu2 struct {
	BadThing string
}

func (s2 *Stu2) OnTeacherComing() {
	fmt.Printf("stu2 Stoping doing %s\n", s2.BadThing)
}

type Stu3 struct {
	BadThing string
}

func (s3 *Stu3) OnTeacherComing() {
	fmt.Printf("stu3 Stoping doing %s\n", s3.BadThing)
}

type ClassMonitor struct {
	listenerList []Listener
}

func (cm *ClassMonitor) AddListener(listener Listener) {
	cm.listenerList = append(cm.listenerList, listener)
}

func (cm *ClassMonitor) RemoveListener(listener Listener) {
	for i, l := range cm.listenerList {
		if l == listener {
			cm.listenerList = append(cm.listenerList[:i], cm.listenerList[i+1:]...)
		}
	}
}

func (cm *ClassMonitor) Notify() {
	for _, listener := range cm.listenerList {
		listener.OnTeacherComing()
	}
}

// === 逻辑层 ===
func main() {
	s1 := &Stu1{
		"抄作业",
	}
	s2 := &Stu2{
		"玩王者荣耀",
	}
	s3 := &Stu3{
		"看别人玩王者荣耀",
	}
	cm := new(ClassMonitor)
	cm.AddListener(s1)
	cm.AddListener(s2)
	cm.AddListener(s3)

	fmt.Println("Teacher is coming!!!")
	cm.Notify()
}
