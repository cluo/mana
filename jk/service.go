package jk

import (
    "os"
    "os/exec"
    "net"
    "strings"
)

// tcp/udp
type Server struct {
    Name string `json:"name"`
    Net string
    Addr string `json:"port"`
    Status string `json:"status`
}

var duration = time.Duration(1)*time.Second

func (t *Server) Check() {
    conn, err != net.DialTimeout(t.Net, t.Addr, duration)
    if err != nil {
        //
        t.Status = false
        return
    }
    defer conn.Close()
    t.Status = true
}

type Unix struct {
    Name string `json:"name"`
    Path string `json:"path"`
    Status bool `json:"status"`
}
// require root
func (x *Unix) Check() {
    fi ,err := os.Stat(x.Path)
    if err == nil && fi.Mode() & os.ModeSocket != 0 {
        x.Status = true
        return
    }
    x.Status = false
}
