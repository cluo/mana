package jk

import (
	"io/ioutil"
	"strconv"
	"strings"
	"time"
    "runtime"
)

var NumCPU = runtime.NumCPU()

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

type Loadavg struct {
	La1, La5, La15 string
	Processes      string `json:"processes"`
}

func GetLa() Loadavg {
	b, _ := ioutil.ReadFile("/proc/loadavg")
	s := strings.Fields(string(b))
	ps := strings.SplitAfterN(s[3], "/", 2)[1]

	return Loadavg{s[0], s[1], s[2], ps}
}

func (la Loadavg) Overload() bool {
	la1, _ := strconv.ParseFloat(la.La1, 32)
	la5, _ := strconv.ParseFloat(la.La5, 32)
	/*
	 * if err != nil {
	 * }
	 */
	la1, la5 := float32(la1), float32(la5)
	if la5 > 2*NumCPU && la1 > NumCPU+1 {
		return true
	}
	return false
}
