package info

import (
	"os/exec"
    "mana/cfg"
)
// 自定义的进程检查,
type Process cfg.Process

func (p *Process) String() string {
	s := p.Name + ": " + p.Pid
	return s
}

func (p *Process) Check() error {
	cmd := cf.Base + "/bin/" + p.Name
	out, err := exec.Command(cmd).Output()
	if err != nil {
		return error
	}
	p.Pid = string(out)
    return nil
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
