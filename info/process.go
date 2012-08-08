package info

import (
	"errors"
	"os/exec"
	"path/filepath"
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
	var reply = new(Process)
	reply.Name = name
	var cmd = "./bin/" + name
	pid_b, err := exec.Command(cmd).Output()
	if err != nil {
		a.Log.Println("process", cmd, err)
		return nil, err
	}
	reply.Pid = string(pid_b)
	return reply, nil
}

// CPU或者MEM使用率最高的进程
type TopProcess struct {
	Sort   string
	Num    string
	Result string
}

func (top *TopProcess) String() string {
	s := "Top of processes " + top.Sort + ", " + top.Num + ":\n" + top.Result
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
	var reply = new(TopProcess)
	reply.Sort = sort
	reply.Num = n
	ctop, err := exec.Command(cmd, n).Output()
	if err != nil {
		a.Log.Println("top", cmd, err)
		return nil, err
	}
	reply.Result = string(ctop)
	return reply, nil
}

// 方便以后使用shell脚本获取特定信息
type Shell struct {
	Name, Path string
	Result     string
}

// name如果为""，使用filepath.Base(path)
func (a *Agent) Shell(name, path string) (*Shell, error) {
	if base := filepath.Base(path); base != "." {
		if name == "" {
			name = base
		}
		bts, err := exec.Command(path).Output()
		if err != nil {
			a.Log.Println("Shell", err)
			return nil, err
		}
		return &Shell{name, path, string(bts)}, nil
	}
	return nil, errors.New("Shell path error")
}

func (sh *Shell) String() string {
	s := sh.Name + ":\n" + sh.Result
	return s
}
