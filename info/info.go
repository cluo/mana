/*
   sensors命令:"lm-sensors"
   mpstat, iostat命令:"sysstat"
   注:单词来自词典
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

var Timestr = "2006-01-02 15:04:05 -0700 MST"

type System struct {
	Hostname *Hostname
	Load     *Load
	Temp     *Temp
}

func (s *System) String() string {
	s.Load.Free = s.Load.Free.Format()
	res := s.Hostname.String() + "\n" + s.Load.String() + "\n" + s.Temp.String()
	return res
}

type Agent struct {
	Log *log.Logger
}

func NewAgent(ll *log.Logger) *Agent {
	return &Agent{ll}
}

var DefaultAgent = NewAgent(log.New(os.Stderr, "", log.LstdFlags))

func (a *Agent) System() (*System, error) {
	sys := new(System)
	host, err := a.Hostname()
	if err != nil {
		return nil, errors.New("func Hostname() failed")
	}
	sys.Hostname = host
	load, err := a.Load()
	if err != nil {
		return nil, errors.New("func Load() failed")
	}
	sys.Load = load
	temp, err := a.Temp()
	if err != nil {
		return nil, errors.New("func Temp() failed")
	}
	sys.Temp = temp
	return sys, nil
}

func TopCpu() (*TopProcess, error) {
	return DefaultAgent.Top("10", "cpu")
}

func TopMem() (*TopProcess, error) {
	return DefaultAgent.Top("10", "cpu")
}
