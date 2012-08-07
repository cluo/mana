package info

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

var numcpu = runtime.NumCPU()

// CPU使用率
type Pcpu struct {
	Us float64
	Sy float64
	Wa float64
	Id float64
}

func (p Pcpu) String() string {
	return fmt.Sprintf("us\tsy\twa\tidle\n%.2f\t%.2f\t%.2f\t%.2f", p.Us, p.Sy, p.Wa, p.Id)
}

func GetPcpus() ([]Pcpu, error) {
	var pcpus []Pcpu
	all, err := exec.Command("/usr/bin/mpstat", "-P", "ALL").Output()
	if err != nil {
		elog.Println("/usr/bin/mpstat -P ALL", err)
		return nil, err
	}
	s := strings.SplitAfter(string(all), "\n")
	for i := 3; i <= numcpu+3; i++ {
		var cpu = s[i]
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
		pcpus = append(pcpus, Pcpu{us, sy, wa, id})
	}
	return pcpus, nil
}

type Iostat string

func GetIostat() (Iostat, error) {
	out, err := exec.Command("iostat", "-kdx").Output()
	if err != nil {
		elog.Println("iostat -kdx:", err)
		return "", errors.New("iostat command error")
	}
	return Iostat(out), nil
}

type Loadavg struct {
	La1, La5, La15 string
	Processes      string
}

func (l *Loadavg) String() string {
	return fmt.Sprintf("%s %s %s\t%s", l.La1, l.La5, l.La15, l.Processes)
}

// 系统负载情况
func GetLoadavg() (*Loadavg, error) {
	b, err := ioutil.ReadFile("/proc/loadavg")
	if err != nil {
		elog.Println("reading /proc/loadavg:", err)
		return nil, err
	}
	s := strings.Fields(string(b))

	return &Loadavg{s[0], s[1], s[2], s[3]}, nil
}

func (la *Loadavg) Overload() bool {
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
type Mem struct {
	Total   string
	Used    string
	Free    string
	Buffers string
	Cached  string
}

// 交换分区
type Swap struct {
	Total string
	Used  string
	Free  string
}

// free -o
type Free struct {
	Mem  Mem
	Swap Swap
}

func (f *Free) String() string {
	mem, swap := f.Mem, f.Swap
	s := fmt.Sprintf("Mem:\t%s\t%s\t%s\t%s\t%s\nSwap:\t%s\t%s\t%s",
		mem.Total, mem.Used, mem.Free, mem.Buffers, mem.Cached,
		swap.Total, swap.Used, swap.Free)
	return s
}

func GetFree() (*Free, error) {
	var free = new(Free)
	bts, err := exec.Command("free", "-o", "-b").Output()
	if err != nil {
		elog.Println("free -ob:", err)
		return nil, err
	}
	lines := strings.Split(string(bts), "\n")
	m, s := strings.Fields(lines[1]), strings.Fields(lines[2])

	free.Mem = Mem{Total: m[1], Used: m[2], Free: m[3], Buffers: m[5], Cached: m[6]}
	free.Swap = Swap{Total: s[1], Used: s[2], Free: s[3]}
	return free, nil
}

// 未使用的内存加上缓存 
func (m *Free) Real() float32 {
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
func (m *Free) Format() *Free {
	var r, s = m.Mem, m.Swap

	fr := Mem{
		format(r.Total).String(),
		format(r.Used).String(),
		format(r.Free).String(),
		format(r.Buffers).String(),
		format(r.Cached).String(),
	}
	fs := Swap{
		format(s.Total).String(),
		format(s.Used).String(),
		format(s.Free).String(),
	}
	return &Free{fr, fs}
}

type Load struct {
	Cpu  []Pcpu
	Free *Free
	Load *Loadavg
	IO   Iostat
}

func (l *Load) String() string {
	head := "System Load status\n"
	var cpu string
	for k, v := range l.Cpu {
		if k == 0 {
			cpu += fmt.Sprintf("CPU ALL:\n%s\n", v)
			continue
		}
		cpu += fmt.Sprintf("CPU %d:\n%s\n", k-1, v)
	}
	return fmt.Sprintf("%s\nCPU status: %s\nMemory status:\n%s\n\nLoadavg: %s\n\nIostat:\n%s",
		head, cpu, l.Free, l.Load, l.IO)
}

func GetLoad() (*Load, error) {
	pcpu, err := GetPcpus()
	if err != nil {
		return nil, errors.New("func GetPcpu() failed")
	}
	free, err := GetFree()
	if err != nil {
		return nil, errors.New("func GetFree() failed")
	}
	load, err := GetLoadavg()
	if err != nil {
		return nil, errors.New("func GetLoadavg() failed")
	}
	iostat, err := GetIostat()
	if err != nil {
		return nil, errors.New("func GetIostat() failed")
	}
	return &Load{pcpu, free, load, iostat}, nil
}
