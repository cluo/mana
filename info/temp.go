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
		return
	}
	defer conn.Close()

	bts, err := ioutil.ReadAll(conn)
	if err != nil {
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

type Sensors string

func GetSensors() (Sensors, error) {
	out, err := exec.Command("sensors").Output()
	return Sensors(out), err
}

func (s Sensors) Check() (bool, error) {
	sensor := cf.Base + "/bin/" + "sensors"
	out, err := exec.Command(sensor, string(s)).Output()
	if err != nil {
		return false, err
	}
	if len(out) > 0 {
		return true, nil
	}
	return false, nil
}
