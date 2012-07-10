package config

type Service struct {
    Name string
    Net, Addr string
    Status bool `json:"-"`
    Interval float64
}

func (s *Service) SetTime(t float64) {
    s.Interval = t
}

type Process struct {
    Name string
    Pid string `json:",omitempty"`
    Interval float64
}

func (p *Process) SetTime(t float64) {
    p.Interval = t
}
