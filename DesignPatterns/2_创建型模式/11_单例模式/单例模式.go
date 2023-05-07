package main

import (
	"fmt"
	"sync"
)

//1. 一个类只有一个实例
//2. 自行创建这个实例
//3. 自行向整个系统提供这个唯一实例

// 首字母小写, 保证该类非公有化，使其他包无法直接创建该类的对象值
type singleton struct {
}

// 饿汉式单例，在执行main之前就已经被创造，无论是否会被使用
// 保证指针始终指向唯一实例对象，首字母小写(其他包无法访问)
//var instance = new(singleton)

// GetInstance 为了满足第三点，对外提供一个方法来获取该对象, 只有读权限，无写权限
//饿汉式单例
//func GetInstance() *singleton {
//	return instance
//}

// 懒汉式单例，只有在需要时才被创建
var instance *singleton

// 为了保证高并发运行GetInstance时线程安全，只有一个唯一实例，加入互斥锁, 但是加锁解锁速度较慢，影响性能, 可以使用原子操作, 也可以使用once
//var lock sync.Mutex

// 原子操作
//var initialized uint32

// once
var once sync.Once

func GetInstance() *singleton {
	//标识位判断, 除第一次调用外，其他调用直接在此返回, 不会触发加锁解锁
	//if atomic.LoadUint32(&initialized) == 1 {
	//	return instance
	//}
	//lock.Lock()
	//defer lock.Unlock()
	//if instance == nil {
	//	atomic.StoreUint32(&initialized, 1)
	//	instance = new(singleton)
	//}
	once.Do(func() {
		instance = new(singleton)
	})
	return instance
}

func (s *singleton) Something() {
	fmt.Println("A Func for singleton")
}

func main() {
	s1 := GetInstance()
	s1.Something()
	s2 := GetInstance()
	if s1 == s2 {
		fmt.Println("s1 and s2 point the same address")
	}
}
