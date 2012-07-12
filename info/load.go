package info

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strconv"
	"strings"
)

// CPU使用率
type Pcpu struct {
	Us float64
	Sy float64
	Wa float64
	Id float64
}

func GetPcpu() (*Pcpu, error) {
	out, err := exec.Command("/usr/bin/mpstat").Output()
	if err != nil {
		return nil, err
	}
	s := strings.SplitAfter(string(out), "\n")
	var cpu = s[len(s)-1]
	cur := strings.Fields(cpu)
	/*
	 *us,ni,sy := cur[2],cur[3],cur[4]
	 *wa,hi,si := cur[5],cur[6],cur[7]
	 *st,id := cur[8],cur[10]
	 */
	us, _ := strconv.ParseFloat(cur[2], 64)
	sy, _ := strconv.ParseFloat(cur[4], 64)
	wa, _ := strconv.ParseFloat(cur[5], 64)
	id, _ := strconv.ParseFloat(cur[10], 64)
	return &Pcpu{us, sy, wa, id}, nil
}

type Iostat string

func GetIostat() (Iostat, error) {
	out, err := exec.Command("iostat", "-kdx").Output()
	if err != nil {
		return "", nil
	}
	return Iostat(out), nil
}

type Loadavg struct {
	La1, La5, La15 string
	Processes      string
}

// 系统负载情况
func GetLoadavg() (*Loadavg, error) {
	b, err := ioutil.ReadFile("/proc/loadavg")
	if err != nil {
		return nil, err
	}
	s := strings.Fields(string(b))
	ps := strings.SplitAfterN(s[3], "/", 2)[1]

	return &Loadavg{s[0], s[1], s[2], ps}, nil
}

func (la *Loadavg) Check() bool {
	n := float64(numcpu)
	la1, _ := strconv.ParseFloat(la.La1, 32)
	la5, _ := strconv.ParseFloat(la.La5, 32)
	if la5 > 2*n && la1 > n+1 {
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

// 物理内存
type Realm struct {
	Total   string
	Used    string
	Free    string
	Buffers string
	Cached  string
}

// 交换分区
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

func GetMemory() (memory *Memory, err error) {
	bts, err := exec.Command("free", "-o", "-b").Output()
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(bts), "\n")
	m, s := strings.Fields(lines[1]), strings.Fields(lines[2])

	memory.Mem = Realm{m[1], m[2], m[3], m[5], m[6]}
	memory.Swap = Swapd{s[1], s[2], s[3]}
	return memory, nil
}

// 未使用的内存加上缓存 
func (m *Memory) Real() float32 {
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

// 格式化为易读的数据
func (m *Memory) Format() *Memory {
	var r Realm = m.Mem
	var s Swapd = m.Swap

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
	return &Memory{fr, fs}
}
