package info

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type Adapter struct {
	Name     string
	Receive  int64
	Transmit int64
	Time     time.Time
}

func (a Adapter) String() string {
	return fmt.Sprintf("%s receive:%d, transmit:%d, %s", a.Name, a.Receive, a.Transmit, a.Time)
}

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
		if len(lines[i]) > 10 {
			t := strings.Fields(lines[i])
			name := strings.Trim(t[0], ":")
			recv, _ := strconv.ParseInt(t[1], 10, 64)
			tran, _ := strconv.ParseInt(t[10], 10, 64)
			adapters = append(adapters, Adapter{name, recv, tran, date})
		}
	}
	return adapters, nil
}
