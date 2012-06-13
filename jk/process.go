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
var ExecPath = "/home/link0x/src/mana/jk/custom/"

func (p *Process) GetPid() {
    cmd :=  ExecPath + p.Name
    out,err := exec.Command(cmd).Output()
    var pid string
    if err != nil {
        pid = "NO"
    } else {
        pid = string(out)
    }
    p.Pid = pid
}
