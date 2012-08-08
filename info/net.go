package info

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

// 上传和下载的流量，从系统启动累加 
type Traffic struct {
	Name     string
	Receive  float64
	Transmit float64
	Time     time.Time
}

func (t Traffic) String() string {
	recv, trans := ByteSize(t.Receive), ByteSize(t.Transmit)
	ts := t.Time.Format(Timestr)
	return fmt.Sprintf("%s receive:%s, transmit:%s, %v",
		t.Name, recv, trans, ts)
}

// 读取/proc/net/dev
func (a *Agent) Traffic() ([]Traffic, error) {
	bts, err := ioutil.ReadFile("/proc/net/dev")
	if err != nil {
		a.Log.Println("Adapter", err)
		return nil, err
	}
	date := time.Now()
	lines := strings.Split(string(bts), "\n")
	var tra []Traffic
	for i := 3; i < len(lines); i++ {
		t := strings.Fields(lines[i])
		if len(t) == 17 {
			name := strings.Trim(t[0], ":")
			recv, _ := strconv.ParseFloat(t[1], 64)
			tran, _ := strconv.ParseFloat(t[10], 64)
			tra = append(tra, Traffic{name, recv, tran, date})
		}
	}
	return tra, nil
}
