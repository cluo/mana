package info

import (
	"fmt"
	"net"
	"time"
)

// tcp/udp 系统服务
type Service struct {
	Name    string
	Address string
	Status  bool
}

type Args_s struct {
	Name      string
	Net, Addr string
	Timeout   time.Duration
}

func (srv *Service) Get(args *Args_s, reply *Service) error {
	reply.Name = args.Name
	reply.Address = args.Net + "://" + args.Addr
	var timeout = 1 * time.Second
	if args.Timeout > 1*time.Second {
		timeout = args.Timeout
	}
	conn, err := net.DialTimeout(args.Net, args.Addr, timeout)
	if err != nil {
		ErrorLog.Println("net.Dial,", args.Net, args.Addr)
		return err
	}
	conn.Close()
	reply.Status = true
	return nil
}

func (srv *Service) String() string {
	return fmt.Sprintf("%s:%s:%s", srv.Name, srv.Address, srv.Status)
}
