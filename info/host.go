package info

import (
	"io/ioutil"
	"os"
	"time"
)

type Host struct {
	Hostname string
	Boot     time.Time
	Uptime   string
}

func GetHost() (h *Host, err error) {
	b, err := ioutil.ReadFile("/proc/uptime")
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(b); i++ {
		if b[i] == ' ' {
			b = b[0:i]
			break
		}
	}
	t := string(b) + "s"
	d, err := time.ParseDuration(t)
	if err != nil {
		return nil, err
	}
	boot := time.Now().Add(-d)

	hn, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	return &Host{hn, boot, d.String()}, nil
}
