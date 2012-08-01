package info

import (
	"fmt"
	"net"
	"os/exec"
)

// tcp/udp 系统服务
type Tcp struct {
	Name    string
	Address string
	Status  bool
}

func (a *Agent) Tcp(name, addr, port string) (*Tcp, error) {
	tcp := new(Tcp)
	tcp.Name = name
	tcp.Address = addr + ":" + port
	conn, err := net.Dial("tcp", tcp.Address)
	if err != nil {
		elog.Println("service tcp:", name, err)
		return nil, err
	}
	defer conn.Close()
	tcp.Status = true
	return tcp, nil
}

func (t *Tcp) String() string {
	return fmt.Sprintf("%s:%s:%s", t.Name, t.Address, t.Status)
}

type Udp struct {
	Name    string
	Address string
	Status  bool
}

func (a *Agent) Udp(name, addr, port string) (*Udp, error) {
	udp := new(Udp)
	udp.Name = name
	udp.Address = addr + ":" + port
	cmd := exec.Command("nc", "-zvu", addr, port)
	_, err := cmd.Output()
	if err != nil {
		elog.Println("service udp:", name, "error")
		return nil, err
	}
	udp.Status = true
	return udp, nil
}

func (u *Udp) String() string {
	return fmt.Sprintf("%s:%s:%s", u.Name, u.Address, u.Status)
}
