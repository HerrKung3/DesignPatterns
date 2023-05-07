package main

import "fmt"

//如果使用继承，会导致父类的任何变换都可能影响到子类的行为。如果使用对象组合，就降低了这种依赖关系，优先使用组合

type Cat struct {
}

func (c *Cat) Eat() {
	fmt.Println("Cat is eating")
}

// CatB 添加睡觉的能力(使用继承)
type CatB struct {
	Cat
}

func (cb *CatB) Sleep() {
	fmt.Println("Cat is sleeping")
}

// CatC 添加睡觉的能力(使用组合)
type CatC struct {
}

func (cc *CatC) Eat(c *Cat) {
	c.Eat()
}

func (cc *CatC) Sleep() {
	fmt.Println("Cat is sleeping")
}

func main() {
	c := &Cat{}
	cb := &CatB{}
	cb.Eat()
	cb.Sleep()
	cc := &CatC{}
	cc.Sleep()
	cc.Eat(c)
}
