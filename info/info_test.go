package info

import (
	"fmt"
	"testing"
)

var agent = DefaultAgent

func TestHostname(t *testing.T) {
	h, err := agent.Hostname()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(h)
}

func TestPcpu(t *testing.T) {
	pcpu, err := agent.Pcpu()
	if err != nil {
		t.Error(err)
	}
	for k, v := range pcpu {
		if k == 0 {
			fmt.Printf("cpu ALL:\n%s\n", v)
		} else {
			fmt.Printf("cpu %d:\n%s\n", k-1, v)
		}
	}
}

func TestLoadavg(t *testing.T) {
	loadavg, err := agent.Loadavg()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(loadavg)
}

func TestFree(t *testing.T) {
	free, err := agent.Free()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(free)
}

func TestHddtemp(t *testing.T) {
	hts, err := agent.Hddtemp()
	if err != nil {
		t.Error(err)
	}
	for k, v := range hts {
		fmt.Println(k, v)
	}
}

func TestProcess(t *testing.T) {
	var name = "nginx_master"
	proc, err := agent.Process(name)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(proc)
}

func TestShell(t *testing.T) {
	var name = "ports listening "
	sh, err := agent.Shell(name, "./bin/netstat")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(sh)
}

func TestTcp(t *testing.T) {
	tcp, err := agent.Tcp("hddtemp", "127.0.0.1", "7634")
	if err != nil {
		t.Log(err)
	}
	fmt.Println(tcp)
}
func TestUdp(t *testing.T) {
	udp, err := agent.Udp("snmp", "127.0.0.1", "161")
	if err != nil {
		t.Log(err)
	}
	fmt.Println(udp)
}

func TestTraffic(t *testing.T) {
	as, err := agent.Traffic()
	if err != nil {
		t.Error(err)
	}
	for _, v := range as {
		fmt.Println(v)
	}
}

func TestTop(t *testing.T) {
	top, err := agent.Top("10", "cpu")
	if err != nil {
		t.Log("top sorted by cpu", err)
	}
	fmt.Println(top)
	top, err = agent.Top("10", "mem")
	if err != nil {
		t.Log("top sorted by mem", err)
	}
	fmt.Println(top)
}

func TestAgentSysTem(t *testing.T) {
	sys, err := agent.System()
	if err != nil {
		t.Log(err)
	}
	fmt.Printf("%s\n", sys)
}
