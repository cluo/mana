package jk

import (
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type Uptime struct {
	Hostname string
	Boot     time.Time
	Duration string
}

func newUptime(h, d string, b time.Time) *Uptime {
	return &Uptime{h, b, d}
}

func GetUptime() (u *Uptime, err error) {
	b, err := ioutil.ReadFile("/proc/uptime")
	if err != nil {
		return nil, err
	}

	s := strings.Fields(string(b))[0] + "s"
	if err != nil {
		return nil, err
	}
	d, err := time.ParseDuration(s)
	if err != nil {
		return nil, err
	}
	boot := time.Now().Add(-d)

	h, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	return newUptime(h, d.String(), boot), nil
}
