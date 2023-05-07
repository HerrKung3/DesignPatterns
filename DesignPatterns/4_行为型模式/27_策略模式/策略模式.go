package main

import "fmt"

//准备一组算法，并将每一个算法封装起来，使它们可以互换

//策略模式优点：
//			1 完美支持开闭原则
//			2 避免多重条件选择语句
//			3 提供了策略复用机制，不同环境类可以方便地复用这些策略
//策略模式缺点：
//			1 客户端必须知道并理解所有策略类，并自行选择策略
//			2 系统将产生许多具体策略类，任何细微变化都会导致新的具体策略类产生

// WeaponStrategy 武器策略(抽象的策略)
type WeaponStrategy interface {
	UseWeapon()
}

//具体策略

type Ak47 struct {
}

func (ak *Ak47) UseWeapon() {
	fmt.Println("Use AK47")
}

type Knife struct {
}

func (k *Knife) UseWeapon() {
	fmt.Println("Use knife")
}

//环境类 发起策略调用

type Hero struct {
	strategy WeaponStrategy
}

// SetWeaponStrategy 设置一个策略方法，给接口赋值
func (h *Hero) SetWeaponStrategy(s WeaponStrategy) {
	h.strategy = s
}

func (h *Hero) Fight() {
	h.strategy.UseWeapon()
}

func main() {
	hero := Hero{}
	hero.SetWeaponStrategy(new(Ak47))
	hero.Fight()
	hero.SetWeaponStrategy(new(Knife))
	hero.Fight()
}
