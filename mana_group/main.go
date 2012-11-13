/*
mana_group 用来收集每个计算机的信息.

就像一开始的目的，就是简单，只要是监控中的服务、进程不区分重要程度，
统一查询时间间隔。我们要做的只是指定group中需要监控那些计算机，至于
服务器需要监控的进程或者服务，
只在agent中设置，这方面具体看mana_agent。
*/
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/smtp"
	"time"
)

var (
	sys_time     = 300 * time.Second
	srv_time     = 29 * time.Second
	proc_time    = 37 * time.Second
	sh_time      = 47 * time.Second
	running_time = 11 * time.Second
)

var (
	help       = flag.Bool("h", false, "help")
	config_dir = flag.String("c", "etc", "config files path")
	mail_warn  = flag.Bool("mail", false, "send mail when status changed")
)

type Group struct {
	Name     string      `json:"name"`
	Computer []*Computer `json:"computer"`
}

func warn_print(mu *MailUser, auth smtp.Auth) {
	for {
		select {
		case warn := <-notify.Warn:
			if *mail_warn {
				go func() {
					addr := mu.Host + ":" + mu.Port
					err := smtp.SendMail(addr, auth, mu.Username,
						mu.To, warn_message(mu, warn))
					if err != nil {
						notify.err <- err
					}
				}()
			}
			fmt.Println(warn)
		case err := <-notify.err:
			fmt.Println(err)
		}
	}
}

func main() {
	if *help {
		flag.PrintDefaults()
		return
	}
	config_group := *config_dir + "/group"
	group_bytes, err := ioutil.ReadFile(config_group)
	if err != nil {
		fmt.Println(err)
		return
	}
	var group Group
	err = json.Unmarshal(group_bytes, &group)
	if err != nil {
		fmt.Println(err)
		return
	}

	var (
		sys     = time.Tick(sys_time)
		srv     = time.Tick(srv_time)
		proc    = time.Tick(proc_time)
		sh      = time.Tick(sh_time)
		running = time.Tick(running_time)
	)
	config_mail := *config_dir + "/mail"
	var mu = NewMailUser(config_mail)
	//重新检查具体信息并判断是否需要产生提醒
	go check_if_warn(notify.Retry, notify.Warn)
	go warn_print(mu, smtp_auth(mu))
	//组成员检查
	for {
		select {
		case <-running:
			for _, co := range group.Computer {
				co.status()
			}
		case <-sys:
			for _, co := range group.Computer {
				co.System()
			}
		case <-srv:
			for _, co := range group.Computer {
				co.Tcp()
				co.Udp()
			}
		case <-proc:
			for _, co := range group.Computer {
				co.Process()
			}
		case <-sh:
			for _, co := range group.Computer {
				co.Shell()
			}
		}
	}
}
