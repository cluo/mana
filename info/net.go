package info

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

// 上传和下载的流量，从系统启动累加 
type Adapter struct {
	Name     string
	Receive  float64
	Transmit float64
	Time     time.Time
}

func (a Adapter) String() string {
	recv, trans := ByteSize(a.Receive), ByteSize(a.Transmit)
	return fmt.Sprintf("%s receive:%s, transmit:%s, %v",
		a.Name, recv, trans, a.Time)
}

// 读取/proc/net/dev
func GetAdapters() ([]Adapter, error) {
	bts, err := ioutil.ReadFile("/proc/net/dev")
	if err != nil {
		elog.Println("GetAdapter", err)
		return nil, err
	}
	date := time.Now()
	lines := strings.Split(string(bts), "\n")
	var adapters []Adapter
	for i := 3; i < len(lines); i++ {
		if len(lines[i]) == 16 {
			t := strings.Fields(lines[i])
			name := strings.Trim(t[0], ":")
			recv, _ := strconv.ParseFloat(t[1], 64)
			tran, _ := strconv.ParseFloat(t[10], 64)
			adapters = append(adapters, Adapter{name, recv, tran, date})
		}
	}
	return adapters, nil
}
