package main

import (
	"fmt"
	"io/ioutil"
	"json"
	"mana/info"
	"net/http"
	"time"
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

type Notify struct {
	check chan string
	warn  chan string
	err   chan error
}

func NewNotify() *Notify {
	var redo, warn, err = make(chan string), make(chan string), make(chan error)
	return &Notify{redo, warn, err}
}

var notify = NewNotify()

func readResponse(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		notify.err <- fmt.Errorf("Server(%s) error: %s", co.Address, err)
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
	resp, err := http.Head(co.URI())
	var status bool
	if err != nil {
		notify.err <- err
	} else {
		status = true
	}
	if co.status != status {
		notify.redo <- co.URI()
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

func fromjson(b, v interface{}) error {
	if b != nil {
		err := json.Unmarshal(b, v)
		if err != nil {
			notify.err <- fmt.Errorf("Data (%s) error: %s", urL, err)
			return err
		}
	}
	return fmt.Errorf("Data []byte null")
}

func (co *Computer) System() {
	if !co.status {
		return
	}
	urL := co.URI() + "/system"
	b := readResponse(url)
	var sy info.System
	err := fromjson(b, sy)
	if err != nil {
		return
	}
	co.sys = sy
}

func (co *Computer) Tcp() {
	if !co.status {
		return
	}
	// tcp check
	urL := co.URI() + "/service?q=tcp&name=" + name
	b := readResponse(urL)
	var tcp []*info.Service
	err := fromjson(b, &tcp)
	if err != nil {
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
	urL := co.URI() + "/service?q=udp&name=" + name
	b := readResponse(urL)
	var udp []*info.Service
	err := fromjson(b, &udp)
	if err != nil {
		return
	}
	check_service_status(co.URI(), udp, co.udp)
	co.udp = udp
}

func (co *Computer) Process(name string) {
	if !co.status {
		return
	}
	urL := co.URI() + "/process?q=" + name
	b := readResponse(urL)
	var pr []*info.Process
	err := fromjson(b, &pr)
	if err != nil {
		return
	}
	check_process_status(co.URI, pr, co.proc)
	co.proc = pr
}

func (co *Computer) Shell(name string) {
	if !co.status {
		return
	}
	urL := co.URI() + "/custom?q=" + name
	b := readResponse(url)
	var sh []*info.Shell
	err := fromjson(b, &sh)
	if err != nil {
		return
	}
	check_shell_status(co.URI(), sh, co.sh)
	co.sh = sh
}
