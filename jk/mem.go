package jk

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

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
