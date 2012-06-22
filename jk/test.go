package main

import (
"os/exec"
"fmt"
"io/ioutil"
)

type Sensors string

func GetSensors() Sensors {
    out, _ := exec.Command("sensors").Output()
    return Sensors(out)
}

func (s Sensors) Check() bool {
    out, _ := exec.Command("./exec/sensors", string(s)).Output()
    if string(out) == "Y" {
        return true
    }
    return false
}

func main() {
    s,err := ioutil.ReadFile("./exec/sensors.txt")
    if err != nil {
        fmt.Println(err)
    }
    sensors := Sensors(s)
    b := sensors.Check()
    fmt.Print(sensors)
    fmt.Println(b)
}
