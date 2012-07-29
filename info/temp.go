package info

import (
	"bytes"
	"io/ioutil"
	"net"
	"os/exec"
	"time"
)

type Hddtemp struct {
	Dev  string
	Desc string
	Temp string
}

func newHddtemp(dev, des, temp string) Hddtemp {
	return Hddtemp{dev, des, temp}
}

func GetHddtemps() (temps []Hddtemp, err error) {
	conn, err := net.DialTimeout("tcp", "127.0.0.1:7634", 2*time.Second)
	if err != nil {
		ErrorLog.Println("tcp://127.0.0.1:7634", err)
		return
	}
	defer conn.Close()

	bts, err := ioutil.ReadAll(conn)
	if err != nil {
		ErrorLog.Println("reading from 127.0.0.1:7634:", err)
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

type Sensor string

func GetSensor() (Sensor, error) {
	out, err := exec.Command("sensors").Output()
	return Sensor(out), err
}

type Temp struct {
	Disk    []Hddtemp
	Sensors Sensor
}

func GetTemp() (*Temp, error) {
	disks, err := GetHddtemps()
	if err != nil {
		return nil, err
	}
	sensors, err := GetSensor()
	if err != nil {
		return nil, err
	}
	return &Temp{disks, sensors}, nil
}
