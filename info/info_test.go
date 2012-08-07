package info

import (
	"fmt"
	"testing"
)

func TestGetHostname(t *testing.T) {
	fmt.Println("-testing GetHost -")
	h, err := GetHostname()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(h)
	fmt.Println("- end -")
}

func TestGetPcpus(t *testing.T) {
	fmt.Println("- testing GetPcpus -")
	pcpu, err := GetPcpus()
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
	fmt.Println("- end -")
}

func TestGetLoadavg(t *testing.T) {
	fmt.Println("- testing GetLoadavg -")
	loadavg, err := GetLoadavg()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(loadavg)
	fmt.Println("- end -")
}

func TestGetFree(t *testing.T) {
	fmt.Println("- testing GetFree -")
	free, err := GetFree()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(free)
	fmt.Println("- end -")
}

func TestGetHddtemps(t *testing.T) {
	fmt.Println("- testing GetHddtemps -")
	hts, err := GetHddtemps()
	if err != nil {
		t.Error(err)
	}
	for k, v := range hts {
		fmt.Println(k, v)
	}
	fmt.Println("- end -")
}

var agent = new(Agent)

func TestProcess(t *testing.T) {
	fmt.Println("- testing Process -")
	var name = "nginx_master"
	proc, err := agent.Process(name)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(proc)
	fmt.Println("- end -")
}

func TestShell(t *testing.T) {
	fmt.Println("- testing shell -")
	var name = "ports listening "
	sh, err := agent.Shell(name, "./bin/netstat")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(sh)
	fmt.Println("- end -")
}

func TestTcp(t *testing.T) {
	fmt.Println("- testing Tcp -")
	tcp, err := agent.Tcp("hddtemp", "127.0.0.1", "7634")
	if err != nil {
		t.Log(err)
	}
	fmt.Println(tcp)
	fmt.Println("- end -")
}
func TestUdp(t *testing.T) {
	fmt.Println("- testing Udp -")
	udp, err := agent.Udp("snmp", "127.0.0.1", "161")
	if err != nil {
		t.Log(err)
	}
	fmt.Println(udp)
	fmt.Println("- end -")
}

func TestAdapter(t *testing.T) {
	fmt.Println("- testing Adapter -")
	as, err := GetAdapters()
	if err != nil {
		t.Error(err)
	}
	for _, v := range as {
		fmt.Println(v)
	}
	fmt.Println("- end -")
}

func TestTop(t *testing.T) {
	fmt.Println("- testing TopProcess -")
	top, err := agent.Top("10", "cpu")
	if err != nil {
		t.Log("top sorted by cpu", err)
	}
	fmt.Println(top)
	fmt.Println("- end -")
	top, err = agent.Top("10", "mem")
	if err != nil {
		t.Log("top sorted by mem", err)
	}
	fmt.Println(top)
	fmt.Println("- end -")
}

func TestAgentSysTem(t *testing.T) {
	fmt.Println("---------------------------")
	sys, err := agent.System()
	if err != nil {
		t.Log(err)
	}
	fmt.Printf("%s\n", sys)
}
