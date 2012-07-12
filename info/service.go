package info

import (
	"mana/cfg"
	"net"
	"time"
)

// tcp/udp 系统服务
type Service cfg.Service

var timeout = 1 * time.Second

func (t *Service) Check() error {
	conn, err := net.DialTimeout(t.Net, t.Addr, timeout)
	if err != nil {
		t.Status = false
		return err
	} else {
		t.Status = true
	}
	conn.Close()
	return nil
}
