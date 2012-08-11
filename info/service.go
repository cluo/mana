package info

import (
	"fmt"
	"net"
	"os/exec"
)

type Service struct {
	Name    string
	Net     string
	Address string
	Status  bool
}

func (s *Service) String() string {
	return fmt.Sprintf("%s:%s:%s:%t", s.Name, s.Net, s.Address, s.Status)
}

// tcp 服务
// name服务名称，addr、port ip地址和端口
func (a *Agent) Tcp(name, addr, port string) *Service {
	address := addr + ":" + port
	var tcp = &Service{Name: name, Net: "tcp", Address: address}
	conn, err := net.Dial("tcp", tcp.Address)
	if err != nil {
		a.Log.Println("service tcp:", name, err)
		return tcp
	}
	conn.Close()
	tcp.Status = true
	return tcp
}

// udp服务
// Udp 只用于检查本机地址,如"127.0.0.1"、"192.168.1.10"等，port端口
func (a *Agent) Udp(name, addr, port string) *Service {
	address := addr + ":" + port
	var udp = &Service{Name: name, Net: "udp", Address: address}
	cmd := exec.Command("nc", "-zvu", addr, port)
	_, err := cmd.Output()
	if err != nil {
		a.Log.Println("service udp:", name, "down")
		return udp
	}
	udp.Status = true
	return udp
}
