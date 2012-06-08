package jk

import (
	"os/exec"
	"strconv"
	"strings"
)

type Pcpu struct {
	Us float64 `json:"%us"`
	Sy float64 `json:"%sy"`
	Id float64 `json:"%id"`
}

func GetPcpu() (Pcpu, string) {
	out, _ := exec.Command("mpstat", "-P", "ALL")
	s := strings.SplitAfter(string(out), "\n")
	var cpu string
	for _, v := range s {
		if strings.Contains(v, "all") {
			cpu = v
			break
		}
	}
	cur := strings.Fields(cpu)
	/*
	 *us,ni,sy := cur[2],cur[3],cur[4]
	 *wa,hi,si := cur[5],cur[6],cur[7]
	 *st,id := cur[8],cur[10]
	 */
	us, _ := strconv.ParseFloat(cur[2], 64)
	sy, _ := strconv.ParseFloat(cur[4], 64)
	id, _ := strconv.ParseFloat(cur[10], 64)
	return Pcpu{us, sy, id}, string(out)
}

type Iostat string

func GetIostat() Iostat{
    out,_ := exec.Command("iostat","-kd").Output()
    return Iostat(out)
}
