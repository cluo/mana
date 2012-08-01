package info

import (
	"os/exec"
)

// 自定义的进程检查,
type Process struct {
	Name string
	Pid  string
}

func (p *Process) String() string {
	s := p.Name + ":" + p.Pid
	return s
}

func (a *Agent) Process(name string) (*Process, error) {
	var reply = new(Process)
	reply.Name = name
	var cmd = "./bin/" + name
	pid_b, err := exec.Command(cmd).Output()
	if err != nil {
		elog.Println("process", cmd, err)
		return nil, err
	}
	reply.Pid = string(pid_b)
	return reply, nil
}

// CPU使用率最高的进程
type TopCpu struct {
	Num    string
	Result string
}

func (a *Agent) TopBycpu(n string) (*TopCpu, error) {
	var reply = new(TopCpu)
	reply.Num = n
	cmd := "./bin/" + "cpu_top"
	ctop, err := exec.Command(cmd, n).Output()
	if err != nil {
		elog.Println("top cpu", cmd, err)
		return nil, err
	}
	reply.Result = string(ctop)
	return reply, nil
}

type TopMem struct {
	Num    string
	Result string
}

func (a *Agent) TopBymem(n string) (*TopMem, error) {
	var reply = new(TopMem)
	reply.Num = n
	cmd := "./bin/" + "mem_top"
	mtop, err := exec.Command(cmd, n).Output()
	if err != nil {
		elog.Println("top mem", cmd, err)
		return nil, err
	}
	reply.Result = string(mtop)
	return reply, nil
}
