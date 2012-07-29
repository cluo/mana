package info

import (
	"os/exec"
)

// 自定义的进程检查,
type Process struct {
	Name string
	Pid  string
}

func (p *Process) Get(name *string, reply *Process) error {
	reply.Name = *name
	var cmd = "../bin/" + *name
	pid_b, err := exec.Command(cmd).Output()
	if err != nil {
		ErrorLog.Println("exec,", cmd, err)
		return err
	}
	reply.Pid = string(pid_b)
	return nil
}

func (p *Process) String() string {
	s := p.Name + ": " + p.Pid
	return s
}

// CPU使用率最高的进程
type TopCpu struct {
	Num    string
	Result string
}

func (top *TopCpu) Get(n *string, reply *TopCpu) error {
	reply.Num = *n
	cmd := "../bin/" + "cpu_top"
	ctop, err := exec.Command(cmd, *n).Output()
	if err != nil {
		ErrorLog.Println("exec,", cmd, err)
		return err
	}
	reply.Result = string(ctop)
	return nil
}

type TopMem struct {
	Num    string
	Result string
}

func (top *TopMem) Get(n *string, reply *TopMem) error {
	reply.Num = *n
	cmd := "../bin/" + "cpu_mem"
	mtop, err := exec.Command(cmd, *n).Output()
	if err != nil {
		ErrorLog.Println("exec,", cmd, err)
		return err
	}
	reply.Result = string(mtop)
	return nil
}
