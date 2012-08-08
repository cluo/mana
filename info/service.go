package info

import (
	"fmt"
	"net"
	"os/exec"
)

// tcp 服务
type Tcp struct {
	Name    string
	Address string
	Status  bool
}

// name服务名称，addr、port ip地址和端口
func (a *Agent) Tcp(name, addr, port string) (*Tcp, error) {
	tcp := new(Tcp)
	tcp.Name = name
	tcp.Address = addr + ":" + port
	conn, err := net.Dial("tcp", tcp.Address)
	if err != nil {
		a.Log.Println("service tcp:", name, err)
		return nil, err
	}
	defer conn.Close()
	tcp.Status = true
	return tcp, nil
}

func (t *Tcp) String() string {
	return fmt.Sprintf("%s:%s:%t", t.Name, t.Address, t.Status)
}

// udp服务
type Udp struct {
	Name    string
	Address string
	Status  bool
}

// Udp 只用于检查本机地址,如"127.0.0.1"、"192.168.1.10"等，port端口
func (a *Agent) Udp(name, addr, port string) (*Udp, error) {
	udp := new(Udp)
	udp.Name = name
	udp.Address = addr + ":" + port
	cmd := exec.Command("nc", "-zvu", addr, port)
	_, err := cmd.Output()
	if err != nil {
		a.Log.Println("service udp:", name, "error")
		return nil, err
	}
	udp.Status = true
	return udp, nil
}

func (u *Udp) String() string {
	return fmt.Sprintf("%s:%s:%t", u.Name, u.Address, u.Status)
}
