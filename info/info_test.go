package info

import (
	"fmt"
	"testing"
)

func TestGetHostname(t *testing.T) {
	fmt.Println("--- testing GetHost ---")
	h, err := GetHostname()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(h)
	fmt.Println("--- end ---")
}

func TestGetPcpu(t *testing.T) {
	fmt.Println("--- testing GetPcpu ---")
	pcpu, err := GetPcpu()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(pcpu)
	fmt.Println("--- end ---")
}

func TestGetHddtemps(t *testing.T) {
	fmt.Println("--- testing GetHddtemps ---")
	hts, err := GetHddtemps()
	if err != nil {
		t.Error(err)
	}
	for k, v := range hts {
		fmt.Println(k, v)
	}
	fmt.Println("--- end ---")
}

func TestGetLoadavg(t *testing.T) {
	fmt.Println("--- testing GetLoadavg ---")
	loadavg, err := GetLoadavg()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(loadavg)
	fmt.Println("--- end ---")
}

func TestGetFree(t *testing.T) {
	fmt.Println("--- testing GetFree ---")
	free, err := GetFree()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(free.Format())
	fmt.Println("--- end ---")
}

var agent = new(Agent)

func TestProcess(t *testing.T) {
	fmt.Println("--- testing Process ---")
	var name = "nginx_master"
	proc, err := agent.Process(name)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(proc)
	fmt.Println("--- end ---")
}

func TestTcp(t *testing.T) {
	fmt.Println("--- testing Tcp ---")
	tcp, err := agent.Tcp("hddtemp", "127.0.0.1", "7634")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(tcp)
	fmt.Println("--- end ---")
}
func TestUdp(t *testing.T) {
	fmt.Println("--- testing Udp ---")
	udp, err := agent.Udp("snmp", "127.0.0.1", "161")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(udp)
	fmt.Println("--- end ---")
}

func TestAdapter(t *testing.T) {
	fmt.Println("--- testing Adapter ---")
	as, err := GetAdapters()
	if err != nil {
		t.Error(err)
	}
	for _, v := range as {
		fmt.Println(v)
	}
	fmt.Println("--- end ---")
}
