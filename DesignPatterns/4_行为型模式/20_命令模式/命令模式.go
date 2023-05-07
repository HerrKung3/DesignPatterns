package main

import "fmt"

//将一个请求封装成一个对象，从而使你可用不同的请求对客户端进行参数化；对请求排队或记录请求日志，以及支持可撤销操作

//通过病例单将病人和医生解耦

type Doctor struct {
}

func (d *Doctor) treatEyes() {
	fmt.Println("Doctor treat eyes")
}

func (d *Doctor) treatNose() {
	fmt.Println("Doctor treat nose")
}

type Command interface {
	Treat()
}

// CommandTreatEyes 治疗眼睛的病历单
type CommandTreatEyes struct {
	doctor *Doctor
}

func (cte *CommandTreatEyes) Treat() {
	cte.doctor.treatEyes()
}

// CommandTreatNose 治疗鼻子的病历单
type CommandTreatNose struct {
	doctor *Doctor
}

func (ctn *CommandTreatNose) Treat() {
	ctn.doctor.treatNose()
}

// Nurse 护士收集病例单，将病例单发送给医生
type Nurse struct {
	CmdList []Command
}

func (n *Nurse) Notify() {
	if n.CmdList == nil {
		return
	}
	for _, cmd := range n.CmdList {
		cmd.Treat()
	}
}

func main() {
	//填写病例单
	cmdEyes := new(CommandTreatEyes)
	cmdNose := new(CommandTreatNose)
	//收集病例单
	nurse := new(Nurse)
	nurse.CmdList = append(nurse.CmdList, cmdEyes)
	nurse.CmdList = append(nurse.CmdList, cmdNose)
	//将病例单发送给医生
	nurse.Notify()
}
