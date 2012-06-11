package main

import (
    "os/exec"
    "fmt"
)

type Process struct {
    Name string `json:"name"`
    Pid string
}
func (p Process) String() string {
    return fmt.Sprintf("%s: %s",p.Name,p.Pid)
}
//var Ps []string
var ExecPath = "/home/link0x/src/mana/jk/custom/"

func (p *Process) Get() {
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

func main() {
    var p = Process{Name:"nginx_master"}
    p.Get()
    fmt.Println(p)
}
