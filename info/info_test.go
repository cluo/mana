package info

import (
	"fmt"
	"testing"
)

func TestGetHost(t *testing.T) {
	fmt.Println("--- testing GetHost ---")
	h, err := GetHost()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(h)
	fmt.Println("--- end ---")
}

func TestGetPcpu(t *testing.T) {
	if _, err := GetPcpu(); err != nil {
		t.Error(err)
	}
}

func TestGetHddtemps(t *testing.T) {
	if _, err := GetHddtemps(); err != nil {
		t.Error(err)
	}
}

func TestGetLoadavg(t *testing.T) {
	if _, err := GetLoadavg(); err != nil {
		t.Error(err)
	}
}

func TestGetMemory(t *testing.T) {
	if _, err := GetMemory(); err != nil {
		t.Error(err)
	}
}

var my = new(Myapp)

func TestProcess(t *testing.T) {
	var proc = new(Process)
	var name = "nginx_master"
	if err := my.Process(&name, proc); err != nil {
		t.Error(name, err)
	}
}

func TestTcp(t *testing.T) {
	var tcp = new(Tcp)
	var args = &Tcpip{"hddtemp", "127.0.0.1", "7634", 2}
	if err := my.Tcp(args, tcp); err != nil {
		t.Error(*args, err)
	}
}
