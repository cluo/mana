/*
   sensors命令:"lm-sensors"
   mpstat, iostat命令:"sysstat"
   注:单词来自词典
*/
package info

import (
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
	Traffic  []Traffic
	Temp     *Temp
}

type ByName interface {
	GetName() string
}

func (s *System) String() string {
	s.Load.Free = s.Load.Free.Format()
	traffic := "system traffic\n"
	for k, v := range s.Traffic {
		traffic += v.String()
		if k < len(s.Traffic)-1 {
			traffic += "\n"
		}
	}
	res := s.Hostname.String() + "\n" + s.Load.String() + "\n" + traffic + "\n" + s.Temp.String()
	return res
}

type Agent struct {
	Log *log.Logger
}

func NewAgent(ll *log.Logger) *Agent {
	return &Agent{ll}
}

var DefaultAgent = NewAgent(log.New(os.Stderr, "", log.LstdFlags))

func (a *Agent) System() *System {
	host, _ := a.Hostname()
	load, _ := a.Load()
	temp, _ := a.Temp()
	traffic, _ := a.Traffic()
	return &System{host, load, traffic, temp}
}

func TopCpu() (*TopProcess, error) {
	return DefaultAgent.Top("10", "cpu")
}

func TopMem() (*TopProcess, error) {
	return DefaultAgent.Top("10", "cpu")
}
