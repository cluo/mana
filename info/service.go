package info

import (
	"net"
	"time"
)

// tcp/udp
type Service struct {
	Name string
	// tcp/udp/unix
	Net string
	// 127.0.0.1:80 /var/run/example.sock
	Addr     string
	Status   bool `json:"-"`
	Interval time.Duration
}

var timeout = 1 * time.Second

func (t *Service) Check() {
	conn, err := net.DialTimeout(t.Net, t.Addr, timeout)
	if err != nil {
		//
		t.Status = false
	} else {
		t.Status = true
	}
	conn.Close()
}
