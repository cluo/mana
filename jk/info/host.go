package jk

import (
	"io/ioutil"
	"strings"
	"time"
)

type Host struct {
	Hostname   string `json:"hostname"`
	Domainname string `json:"domainname"`
}

func GetHost() Host {
	s, _ := ioutil.ReadFile("/proc/sys/kernel/hostname")
	h := strings.TrimSpace(string(s))
	bts, _ = ioutil.ReadFile("/proc/sys/kernel/domainname")
	d := strings.TrimSpace(string(s))

	return Host{h, d}
}

type Uptime struct {
	Boot     time.Time
	Duration string `json:"duration"`
}

func GetUptime() Uptime {
	b, _ := ioutil.ReadFile("/proc/uptime")

	s := strings.Fields(string(b))[0] + "s"
	d, _ := time.ParseDuration(s)
	/*
	 * if err != nil {
	 *
	 * }
	 */
	bt := time.Now().Add(-d)

	return Uptime{bt, d.String()}
}
