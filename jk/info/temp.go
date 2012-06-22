package jk

import (
	"bytes"
	"io/ioutil"
	"net"
	"os/exec"
	"time"
)

type HddTemp struct {
	Dev  string
	Desc string
	Temp string
}

func GetHddtemps() ([]HddTemp, error) {
	conn, err := net.DialTimeout("tcp", "127.0.0.1:7634", 2*time.Second)
	var hddtemp []HddTemp
	if err != nil {
		return hddtemp, err
	}
	defer conn.Close()

	bts, err := ioutil.ReadAll(conn)
	if err != nil {
		return hddtemp, err
	}

	var buf = bytes.NewBuffer(bts)
	line, err := buf.ReadBytes('\n')
	if err != nil {
		return hddtemp, err
	}
	for i := 0; i < len(line); i++ {
		s := bytes.Split(line[i], []byte("|"))
		hddtemp = append(hddtemp, HddTemp{string(s[1]), string(s[2]), string(s[3])})
	}
	return hddtemp, nil
}

type Sensors string

func GetSensors() Sensors {
	out, _ := exec.Command("sensors").Output()
	return Sensors(out)
}

func (s Sensors) Check() bool {
    out, _ := exec.Command("../exec/sensors", string(s).Output()
    if string(out) == "Y" {
        return true
    }
    return false
}
