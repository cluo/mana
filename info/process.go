package info

import (
	"errors"
	"os/exec"
	"strings"
)

// 自定义的进程检查脚本
type Process struct {
	Name string
	Pid  string
}

func (p *Process) String() string {
	s := p.Name + ":" + p.Pid
	return s
}

// 名称=name的脚本放置在bin目录下
func (a *Agent) Process(name string) (*Process, error) {
	if name == "" {
		return nil, errors.New(`Process(name string) name == ""`)
	}
	var p = new(Process)
	p.Name = name
	var cmd = "./bin/" + name
	pid, err := exec.Command(cmd).Output()
	if err != nil {
		a.Log.Println("process", cmd, err)
		return p, err
	}
	p.Pid = strings.TrimSpace(string(pid))
	return p, nil
}

// CPU或者MEM使用率最高的进程
type TopProcess struct {
	Sort   string
	Num    string
	Result string
}

func (top *TopProcess) String() string {
	s := top.Sort + "_" + top.Num + " (rsz: \"kilobytes\", vsz:\"Kib\")->:\n" + top.Result
	return s
}

// n 是一个整数，脚本"bin/mem_top","bin/cpu_top"中取值为5-10
// sort只能是"cpu"或者"mem"
func (a *Agent) Top(n string, sort string) (*TopProcess, error) {
	var cmd string
	if sort == "" || sort == "cpu" {
		sort = "cpu"
		cmd = "./bin/" + "cpu_top"
	} else if sort == "mem" {
		cmd = "./bin/" + "mem_top"
	} else {
		return nil, errors.New(`sort must be "cpu" or "mem"`)
	}
	var p = new(TopProcess)
	p.Sort = sort
	p.Num = n
	top, err := exec.Command(cmd, n).Output()
	if err != nil {
		a.Log.Println("top", cmd, err)
		return p, err
	}
	p.Result = string(top)
	return p, nil
}

// 方便以后使用shell脚本获取特定信息
type Shell struct {
	Name   string
	Result string
}

// name如果为""，使用filepath.Base(path)
func (a *Agent) Shell(name string) (*Shell, error) {
	if name == "" {
		return nil, errors.New(`Shell(name string) name == ""`)
	}
	var cmd = "./bin/" + name
	var sh = &Shell{Name: name}
	b, err := exec.Command(cmd).Output()
	if err != nil {
		a.Log.Println("Shell", err)
		return sh, err
	}
	sh.Result = string(b)
	return sh, nil
}

func (sh *Shell) String() string {
	s := sh.Name + ":\n" + sh.Result
	return s
}
