package jk

import (
	"fmt"
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

func GetIostat() Iostat {
	out, _ := exec.Command("iostat", "-kd").Output()
	return Iostat(out)
}

var NumCPU = runtime.NumCPU()

type Loadavg struct {
	La1, La5, La15 string
	Processes      string `json:"processes"`
}

func GetLa() Loadavg {
	b, _ := ioutil.ReadFile("/proc/loadavg")
	s := strings.Fields(string(b))
	ps := strings.SplitAfterN(s[3], "/", 2)[1]

	return Loadavg{s[0], s[1], s[2], ps}
}

func (la Loadavg) Check() bool {
	la1, _ := strconv.ParseFloat(la.La1, 32)
	la5, _ := strconv.ParseFloat(la.La5, 32)
	/*  
	 * if err != nil {
	 * }
	 */
	la1, la5 := float32(la1), float32(la5)
	if la5 > 2*NumCPU && la1 > NumCPU+1 {
		return true
	}
	return false
}

type ByteSize float64

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
)

func (b ByteSize) String() string {
	switch {
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	}
	return fmt.Sprintf("%.2fB", b)
}

// mem
type Realm struct {
	Total   string
	Used    string
	Free    string
	Buffers string
	Cached  string
}

// swap
type Swapd struct {
	Total string
	Used  string
	Free  string
}

// free -o
type Memory struct {
	Mem  Realm
	Swap Swapd
}

func GetMem() (statm Memory, err error) {
	var bts []byte
	bts, err = exec.Command("free", "-o", "-b").Output()
	if err != nil {
		return
	}
	lines := strings.Split(string(bts), "\n")
	m, s := strings.Fields(lines[1]), strings.Fields(lines[2])

	statm.Mem = Realm{m[1], m[2], m[3], bz[5], bz[6]}
	statm.Swap = Swapd{s[1], s[2], z[3]}
	return
}

func (m Memory) RealFree() float32 {
	r := m.Mem
	free, _ := strconv.ParseFloat(r.Free, 32)
	buffers, _ := strconv.ParseFloat(r.Buffers, 32)
	cached, _ := strconv.ParseFloat(r.Cached, 32)

	rf := free + buffers + cached

	return float32(rf)
}

func format(s string) ByteSize {
	bys, _ := strconv.ParseFloat(s, 64)
	return ByteSize(bys)
}

func (m Memory) Format() Memory {
	var r Realm = m.Realm
	var s Swapd = m.Swapd

	fr := Realm{
		format(r.Total).String(),
		format(r.Used).String(),
		format(r.Free).String(),
		format(r.Buffers).String(),
		format(r.Cached).String(),
	}
	fs := Swapd{
		format(s.Total).String(),
		format(s.Used).String(),
		format(s.Free).String(),
	}
	return Memory{fr, fs}
}
