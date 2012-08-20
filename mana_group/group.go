package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mana/info"
	"net/http"
)

type Computer struct {
	Address string `json:"address"`
	Ignore  bool   `json:"ignore"`
	status  bool
	sys     *info.System
	tcp     []*info.Service
	udp     []*info.Service
	proc    []*info.Process
	sh      []*info.Shell
}

func (co *Computer) URI() string {
	return fmt.Sprintf("http://%s", co.Address)
}

type Retry struct {
	URL    string
	Class  string
	Status bool
}

type Notify struct {
	Retry chan Retry
	Warn  chan string
	err   chan error
}

func NewNotify() *Notify {
	var retry, warn, err = make(chan Retry), make(chan string), make(chan error)
	return &Notify{retry, warn, err}
}

var notify = NewNotify()

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
	if block != nil {
		return aesDecrypt(block, b)
	}
	return b
}

func (co *Computer) Status() {
	_, err := http.Head(co.URI())
	status := false
	if err != nil {
		notify.err <- err
	} else {
		status = true
	}
	if co.status != status {
		notify.Retry <- Retry{co.URI(), "status", status}
	}
	co.status = status
}

func (co *Computer) Stat() string {
	urL := co.Address + "/status"
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
	if !co.status {
		return
	}
	urL := co.URI() + "/system"
	b := readResponse(urL)
	var sy info.System
	err := fromjson(b, &sy)
	if err != nil {
		notify.err <- err
		return
	}
	co.sys = &sy
}

func (co *Computer) Tcp() {
	if !co.status {
		return
	}
	// tcp check
	urL := co.URI() + "/service?q=tcp"
	b := readResponse(urL)
	var tcp []*info.Service
	err := fromjson(b, &tcp)
	if err != nil {
		notify.err <- err
		return
	}
	check_service_status(co.URI(), tcp, co.tcp)
	co.tcp = tcp
}

func (co *Computer) Udp() {
	if !co.status {
		return
	}
	// tcp check
	urL := co.URI() + "/service?q=udp"
	b := readResponse(urL)
	var udp []*info.Service
	err := fromjson(b, &udp)
	if err != nil {
		notify.err <- err
		return
	}
	check_service_status(co.URI(), udp, co.udp)
	co.udp = udp
}

func (co *Computer) Process() {
	if !co.status {
		return
	}
	urL := co.URI() + "/process"
	b := readResponse(urL)
	var pr []*info.Process
	err := fromjson(b, &pr)
	if err != nil {
		notify.err <- err
		return
	}
	check_process_status(co.URI(), pr, co.proc)
	co.proc = pr
}

func (co *Computer) Shell() {
	if !co.status {
		return
	}
	urL := co.URI() + "/custom"
	b := readResponse(urL)
	var sh []*info.Shell
	err := fromjson(b, &sh)
	if err != nil {
		notify.err <- err
		return
	}
	check_shell_status(co.URI(), sh, co.sh)
	co.sh = sh
}
