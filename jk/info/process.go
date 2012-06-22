package jk

import (
    "os/exec"
)

type Process struct {
    Name string `json:"name"`
    Pid string
}
func (p Process) String() string {
    s := p.Name + ": " + p.pid
    return s
}
//var Ps []string
var ExecPath = "../exec/"

func (p *Process) GetPid() {
    cmd :=  ExecPath + p.Name
    out,err := exec.Command(cmd).Output()
    var pid string
    if err != nil {
        pid = "N"
    } else {
        pid = string(out)
    }
    p.Pid = pid
}

func (p *Process) Check() bool {
    if p.Pid == "N" {
        return false
    }
    return true
}

func Top5cpu() string {
    cmd := ExecPath + "top5_cpu"
    out,_ := exec.Command(cmd).Output()
    return string(out)
}

func Top5mem() string {
    cmd := ExecPath + "top5_mem"
    out,_ := exec.Command(cmd).Output()
    return string(out)
}
