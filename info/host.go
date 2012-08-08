package info

import (
	"io/ioutil"
	"os"
	"time"
)

// 服务器主机名、启动时间以及运行时间
type Hostname struct {
	Name   string
	Boot   time.Time
	Uptime string
}

func (h *Hostname) String() string {
	return h.Name + ":" + h.Boot.Format(Timestr) + ":" + h.Uptime
}

// 读取/proc/uptime，第一数值即系统运行时间，单位为秒(s)
func (a *Agent) Hostname() (*Hostname, error) {
	b, err := ioutil.ReadFile("/proc/uptime")
	if err != nil {
		a.Log.Println("ReadFile /proc/uptime:", err)
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
		a.Log.Println("time.ParseDuration:", err)
		return nil, err
	}
	boot := time.Now().Add(-d)

	hostname, err := os.Hostname()
	if err != nil {
		a.Log.Println("hostname:", err)
		return nil, err
	}

	return &Hostname{hostname, boot, d.String()}, nil
}
