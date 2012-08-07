/*
   sensors命令:"lm-sensors"
   mpstat, iostat命令:"sysstat"
*/
package info

import (
	"errors"
	"log"
	"os"
)

func NewLogger(name string) *log.Logger {
	file, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("opening log:", err)
	}
	return log.New(file, "", log.LstdFlags)
}

var elog = log.New(os.Stderr, "", log.LstdFlags)

type Server struct {
	Hostname *Hostname
	Load     *Load
	Temp     *Temp
}

func (s *Server) String() string {
	s.Load.Free = s.Load.Free.Format()
	res := s.Hostname.String() + "\n" + s.Load.String() + "\n" + s.Temp.String()
	return res
}

type Agent struct{}

func (a *Agent) System() (*Server, error) {
	sys := new(Server)
	host, err := GetHostname()
	if err != nil {
		return nil, errors.New("func GetHostname() failed")
	}
	sys.Hostname = host
	load, err := GetLoad()
	if err != nil {
		return nil, errors.New("func GetLoad() failed")
	}
	sys.Load = load
	temp, err := GetTemp()
	if err != nil {
		return nil, errors.New("func GetTemp() failed")
	}
	sys.Temp = temp
	return sys, nil
}
