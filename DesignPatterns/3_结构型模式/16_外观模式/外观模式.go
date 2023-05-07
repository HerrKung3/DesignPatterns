package main

import (
	"fmt"
)

//为子系统中的一组接口提供一个一致的界面，此模式定义一个高层接口，这个接口使得这一子系统更加容易使用

type TV struct {
}

func (t *TV) On() {
	fmt.Println("Turn on TV")
}

func (t *TV) Off() {
	fmt.Println("Turn off TV")
}

type VoiceBox struct {
}

func (v *VoiceBox) On() {
	fmt.Println("Turn on voice box")
}

func (v *VoiceBox) Off() {
	fmt.Println("Turn off voice box")
}

type Light struct {
}

func (l *Light) On() {
	fmt.Println("Turn on light")
}

func (l *Light) Off() {
	fmt.Println("Turn off light")
}

type Xbox struct {
}

func (x *Xbox) On() {
	fmt.Println("Turn on Xbox")
}

func (x *Xbox) Off() {
	fmt.Println("Turn off Xbox")
}

type MicroPhone struct {
}

func (m *MicroPhone) On() {
	fmt.Println("Turn on microphone")
}

func (m *MicroPhone) Off() {
	fmt.Println("Turn off microphone")
}

type HomePlayerFacade struct {
	tv    TV
	vb    VoiceBox
	light Light
	xbox  Xbox
	mp    MicroPhone
}

func (hp *HomePlayerFacade) DoKTV() {
	fmt.Println("KTV mode")
	hp.tv.On()
	hp.vb.On()
	hp.light.Off()
	hp.xbox.Off()
	hp.mp.On()
}

func (hp *HomePlayerFacade) DoGame() {
	fmt.Println("Game mode")
	hp.tv.On()
	hp.light.On()
	hp.mp.On()
	hp.xbox.On()
}

func main() {
	//不用外观模式, Game mode
	tv := new(TV)
	light := new(Light)
	mp := new(MicroPhone)
	xbox := new(Xbox)
	tv.On()
	light.On()
	mp.On()
	xbox.On()

	//外观模式使用
	hpf := new(HomePlayerFacade)
	hpf.DoGame()
}
