package info

import (
	"os/exec"
)

type Process struct {
	Name string
	Pid  string `json:"-"`
}

func (p *Process) String() string {
	s := p.Name + ": " + p.Pid
	return s
}

func (p *Process) Check() {
	cmd := cf.Base + "/bin/" + p.Name
	out, err := exec.Command(cmd).Output()
	if err != nil {
		return
	}
	p.Pid = string(out)
}

// CPU使用率最高的进程
func Topcpu() (string, error) {
	cmd := cf.Base + "/bin/" + "top_cpu"
	out, err := exec.Command(cmd).Output()
	return string(out), err
}

// 内存占用做多的进程
func Topmem() (string, error) {
	cmd := cf.Base + "/bin/" + "top_mem"
	out, err := exec.Command(cmd).Output()
	return string(out), err
}
