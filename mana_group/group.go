package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mana/info"
	"net/http"
)

//组中所需检查的服务器信息，"address"必需给出;
type Computer struct {
	Address string `json:"address"`
	//"ignore"若为true，则跳过system信息查询
	Ignore bool `json:"ignore"`
	Status bool
	sys    *info.System
	tcp    []*info.Service
	udp    []*info.Service
	proc   []*info.Process
	sh     []*info.Shell
}

func (co *Computer) URI() string {
	return fmt.Sprintf("http://%s", co.Address)
}

type Retry struct {
	URL    string
	Class  string
	Status bool
}

//分别是重新检查，状态变更确认及提醒，错误记录
type Notify struct {
	Retry chan Retry
	Warn  chan string
	err   chan error
}

func NewNotify() *Notify {
	var retry, warn, err = make(chan Retry), make(chan string, 10), make(chan error)
	return &Notify{retry, warn, err}
}

var notify = NewNotify()

//通过http协议获取数据
func readResponse(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		notify.err <- fmt.Errorf("Server(%s) error: %s", url, err)
		return nil
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		notify.err <- fmt.Errorf("Get(%s) error: %s", url, err)
		return nil
	}
	if resp.Header.Get("Mana-Status") != "OK" {
		notify.err <- fmt.Errorf("Get(%s) error: Mana-Status must be \"Ok\"", url)
		return nil
	}
	//aes解密
	if block != nil {
		return aesDecrypt(block, b)
	}
	return b
}

func (co *Computer) status() {
	_, err := http.Head(co.URI())
	status := false
	if err == nil {
		status = true
	}
	if co.Status != status {
		notify.Retry <- Retry{co.URI(), "status", status}
	}
	co.Status = status
}

func (co *Computer) Stat() string {
	urL := co.Address + "/stat"
	b := readResponse(urL)
	if b != nil {
		return string(b)
	}
	return ""
}

func fromjson(b []byte, v interface{}) error {
	if b != nil {
		err := json.Unmarshal(b, v)
		if err != nil {
			notify.err <- fmt.Errorf("Data error: %s", err)
			return err
		}
		return nil
	}
	return fmt.Errorf("Data null")
}

func (co *Computer) System() {
	if !co.Status || co.Ignore {
		return
	}
	urL := co.URI() + "/system"
	b := readResponse(urL)
	var sy info.System
	err := fromjson(b, &sy)
	if err != nil {
		notify.err <- err
		co.sys = nil
		return
	}
	if co.sys != nil {
		check_sys_load(co.Address, sy.Load, co.sys.Load)
		check_sys_temp(co.Address, sy.Temp, co.sys.Temp)
	}
	//check_sys_temp(co.Address, sy.Temp, co.sys.Temp)
	co.sys = &sy
}

func (co *Computer) Tcp() {
	if !co.Status {
		return
	}
	// tcp check
	urL := co.URI() + "/service?q=tcp"
	b := readResponse(urL)
	var tcp []*info.Service
	err := fromjson(b, &tcp)
	if err != nil {
		notify.err <- err
		co.tcp = nil
		return
	}
	check_service_status(co.URI(), tcp, co.tcp)
	co.tcp = tcp
}

func (co *Computer) Udp() {
	if !co.Status {
		return
	}
	// tcp check
	urL := co.URI() + "/service?q=udp"
	b := readResponse(urL)
	var udp []*info.Service
	err := fromjson(b, &udp)
	if err != nil {
		notify.err <- err
		co.udp = nil
		return
	}
	check_service_status(co.URI(), udp, co.udp)
	co.udp = udp
}

func (co *Computer) Process() {
	if !co.Status {
		return
	}
	urL := co.URI() + "/process"
	b := readResponse(urL)
	var pr []*info.Process
	err := fromjson(b, &pr)
	if err != nil {
		notify.err <- err
		co.proc = nil
		return
	}
	check_process_status(co.URI(), pr, co.proc)
	co.proc = pr
}

func (co *Computer) Shell() {
	if !co.Status {
		return
	}
	urL := co.URI() + "/custom"
	b := readResponse(urL)
	var sh []*info.Shell
	err := fromjson(b, &sh)
	if err != nil {
		notify.err <- err
		co.sh = nil
		return
	}
	check_shell_status(co.URI(), sh, co.sh)
	co.sh = sh
}
