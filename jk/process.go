package main

import (
    "os/exec"
    "strings"
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
var ExecPath = "/home/link0x/src/monitor/mana/custom/"

func GetProcess(p string) Process {
    cmd :=  ExecPath + p
    out,_ := exec.Command(cmd).Output()

    proc := strings.Fields(string(out))

    return Process{Name:proc[0],Pid:proc[1]}
}

func main() {
    proc := GetProcess("nginx_master")
    fmt.Println(proc)
}
