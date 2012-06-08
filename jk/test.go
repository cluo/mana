package main

import (
"fmt"
"bytes"
"os/exec"
"strings"
)

type Iostat string

func main() {
    s := "|/dev/sda|VBOX HARDDISK|UNK|*|"
    ss := bytes.Split([]byte(s),[]byte("|"))
    for i := 0; i < len(ss); i++ {
        fmt.Printf("%d: %s\n",i,ss[i])
    }
    out,_ := exec.Command("mpstat","-P","ALL").Output()
    fmt.Println(string(out))

    sss := strings.SplitAfter(string(out),"\n")
    var cpu string
    for _, v := range sss {
        if strings.Contains(v, "all") {
            cpu = v
            break
        }
    }
    cur := strings.Fields(cpu)
    for k,v := range cur {
        fmt.Println(k,v)
    }

    o,_:=exec.Command("iostat","-dk").Output()
    fmt.Println(Iostat(o))
}
