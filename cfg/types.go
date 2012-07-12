package cfg

import "time"

type Service struct {
	Name      string
	Net, Addr string
	Status    bool `json:"-"`
	Interval  time.Duration
}

func (s *Service) Duration(t int64) {
	s.Interval = time.Duration(t) * time.Second
}

type Process struct {
	Name     string
	Pid      string `json:"-"`
	Interval time.Duration
}

func (p *Process) Duration(t int64) {
	p.Interval = time.Duration(t) * time.Second
}
