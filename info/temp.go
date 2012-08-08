package info

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"os/exec"
	"time"
)

// 硬盘温度
type Hddtemp struct {
	Dev  string
	Desc string
	Temp string
}

func (t Hddtemp) String() string {
	return t.Dev + ":" + t.Desc + ":" + t.Temp
}

func newHddtemp(dev, des, temp string) Hddtemp {
	return Hddtemp{dev, des, temp}
}

// 需要hddtemp以守护进程运行，如："sudo hddtemp -d /dev/sd[a-d]"
func (a *Agent) Hddtemp() (temps []Hddtemp, err error) {
	conn, err := net.DialTimeout("tcp", "127.0.0.1:7634", 2*time.Second)
	if err != nil {
		a.Log.Println("tcp://127.0.0.1:7634", err)
		return
	}
	defer conn.Close()

	bts, err := ioutil.ReadAll(conn)
	if err != nil {
		a.Log.Println("reading from 127.0.0.1:7634:", err)
		return
	}

	//var buf = bytes.NewBuffer(bts)
	line := bytes.Split(bts, []byte("\n"))
	for i := 0; i < len(line); i++ {
		s := bytes.Split(line[i], []byte("|"))
		temps = append(temps, newHddtemp(string(s[1]), string(s[2]), string(s[3])))
	}
	return temps, nil
}

// cpu温度
type Sensor string

// 需要命令sensors
func (a *Agent) Sensor() (Sensor, error) {
	out, err := exec.Command("sensors").Output()
	if err != nil {
		a.Log.Println(`sensors error, try to run command "sensors"`)
		return "", errors.New("sensors error")
	}
	return Sensor(out), nil
}

type Temp struct {
	Disks   []Hddtemp
	Sensors Sensor
}

func (temp *Temp) String() string {
	head := "System temperature\n"
	var hdd string
	for _, disk := range temp.Disks {
		hdd += fmt.Sprintf("%s\n", disk)
	}
	return fmt.Sprintf("%s\nHddtemp:\n%s\nSensors:\n%s", head, hdd, temp.Sensors)
}

func (a *Agent) Temp() (*Temp, error) {
	disks, err := a.Hddtemp()
	if err != nil {
		return nil, errors.New("func Hddtemp() failed")
	}
	sensors, err := a.Sensor()
	if err != nil {
		sensors = Sensor("No sensors found")
	}
	/*
		    if err != nil {
				return nil, errors.New("func Sensor() failed")
			}
	*/
	return &Temp{disks, sensors}, nil
}
