package info

import (
	"io/ioutil"
	"os"
	"time"
)

type Hostname struct {
	Name   string
	Boot   time.Time
	Uptime string
}

// 服务器主机名、启动时间以及运行时间:/proc/uptime
func GetHostname() (*Hostname, error) {
	b, err := ioutil.ReadFile("/proc/uptime")
	if err != nil {
		elog.Println("ReadFile /proc/uptime:", err)
		return nil, err
	}
	for i := 0; i < len(b); i++ {
		if b[i] == ' ' {
			b = b[0:i]
			break
		}
	}
	// /proc/uptime
	// first is the uptime of the system (seconds)
	t := string(b) + "s"
	d, err := time.ParseDuration(t)
	if err != nil {
		elog.Println("time.ParseDuration:", err)
		return nil, err
	}
	boot := time.Now().Add(-d)

	hostname, err := os.Hostname()
	if err != nil {
		elog.Println("hostname:", err)
		return nil, err
	}

	return &Hostname{hostname, boot, d.String()}, nil
}
