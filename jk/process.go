package jk

import (
	"os/exec"
    "mana/config"
)

var cf = config.Parse()

type Process struct {
	Name string
	Pid  string
}

func (p Process) String() string {
	s := p.Name + ": " + p.Pid
	return s
}

func Getpid(name string) (p *Process, err error) {
	p.Name = name
	cmd := cf.Basedir + "/bin/" + name
	out, err := exec.Command(cmd).Output()
	if err != nil {
		return nil, err
	}
	p.Pid = string(out)

	return p, nil
}

func Topcpu() (string, error) {
	cmd := cf.Basedir + "/bin/" + "top_cpu"
	out, err := exec.Command(cmd).Output()
	return string(out), err
}

func Topmem() (string, error) {
	cmd := cf.Basedir + "/bin/" + "top_mem"
	out, err := exec.Command(cmd).Output()
	return string(out), err
}
