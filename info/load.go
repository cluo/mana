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
	ID   string
	Us   float64
	Sy   float64
	Wa   float64
	Idle float64
}

func (p Pcpu) String() string {
	return fmt.Sprintf("ID\tus\tsy\twa\tidle\n%s\t%.2f\t%.2f\t%.2f\t%.2f", p.ID, p.Us, p.Sy, p.Wa, p.Idle)
}

// cpu使用率，多cpu的情况，第一个是平均值
func (a *Agent) Pcpu() ([]Pcpu, error) {
	var pcpus []Pcpu
	all, err := exec.Command("/usr/bin/mpstat", "-P", "ALL").Output()
	if err != nil {
		a.Log.Println("/usr/bin/mpstat -P ALL", err)
		return nil, err
	}
	s := strings.SplitAfter(string(all), "\n")
	for i := 3; i <= numcpu+3; i++ {
		var cpu = s[i]
		cur := strings.Fields(cpu)
		/*
			         *ID := cur[1]
					 *us,ni,sy := cur[2],cur[3],cur[4]
					 *wa,hi,si := cur[5],cur[6],cur[7]
					 *st,idle := cur[8],cur[10]
		*/
		us, _ := strconv.ParseFloat(cur[2], 64)
		sy, _ := strconv.ParseFloat(cur[4], 64)
		wa, _ := strconv.ParseFloat(cur[5], 64)
		idle, _ := strconv.ParseFloat(cur[10], 64)
		pcpus = append(pcpus, Pcpu{cur[1], us, sy, wa, idle})
	}
	return pcpus, nil
}

// 获取系统的IO状态
type Iostat string

// 需要命令"iostat"
func (a *Agent) Iostat() (Iostat, error) {
	out, err := exec.Command("iostat", "-kdx").Output()
	if err != nil {
		a.Log.Println("iostat -kdx:", err)
		return "", errors.New("iostat command error")
	}
	str := strings.Replace(string(out), "\n\n", "\n", -1)
	return Iostat(str), nil
}

// 系统负载情况
type Loadavg struct {
	La1, La5, La15 string
	Processes      string
}

func (l *Loadavg) String() string {
	return fmt.Sprintf("%s %s %s\t%s", l.La1, l.La5, l.La15, l.Processes)
}

// 读取/proc/loadavg
func (a *Agent) Loadavg() (*Loadavg, error) {
	b, err := ioutil.ReadFile("/proc/loadavg")
	if err != nil {
		a.Log.Println("reading /proc/loadavg:", err)
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

// ByteSize格式化内存或者流量数据为易读的格式
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

// 使用"free -o -b"命令
func (a *Agent) Free() (*Free, error) {
	var free = new(Free)
	bts, err := exec.Command("free", "-o", "-b").Output()
	if err != nil {
		a.Log.Println("free -ob:", err)
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

// 内存信息转换
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

// Load struct 包含了cpu、内存、系统负载、以及io状态
type Load struct {
	Cpu     []Pcpu
	Free    *Free
	Loadavg *Loadavg
	IO      Iostat
}

func (l *Load) String() string {
	head := "System Load status"
	var cpu string
	for k, v := range l.Cpu {
		if k == 0 {
			cpu += fmt.Sprintf("CPU ALL:\n%s\n", v)
			continue
		}
		cpu += fmt.Sprintf("CPU %d:\n%s\n", k-1, v)
	}
	return fmt.Sprintf("%s\nCPU status: %s\nMemory status:\n%s\n\nLoadavg: %s\n\nIostat:\n%s",
		head, cpu, l.Free, l.Loadavg, l.IO)
}

func (a *Agent) Load() (*Load, error) {
	pcpu, err1 := a.Pcpu()
	free, err2 := a.Free()
	loadavg, err3 := a.Loadavg()
	iostat, err4 := a.Iostat()
	return &Load{pcpu, free, loadavg, iostat}, fmt.Errorf("%s,%s,%s,%s", err1, err2, err3, err4)
}
