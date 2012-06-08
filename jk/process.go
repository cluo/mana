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

func GetProcess(name string) Process {
    cmd :=  ExecPath + name
    out,err := exec.Command(cmd).Output()
    pid := string(out)
    if err != nil {
        pid = "NO"
    }
    return Process{Name:name,Pid:pid}
}

func main() {
    proc := GetProcess("nginx_master")
    fmt.Println(proc)
}
