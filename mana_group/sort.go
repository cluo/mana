package main

import (
	"fmt"
	"mana/info"
	"net/http"
	"sort"
)

//使用sort.Sort包需要实现以下几个方法
type ServiceSlice []*info.Service

func (s ServiceSlice) Len() int           { return len(s) }
func (s ServiceSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ServiceSlice) Less(i, j int) bool { return s[i].GetName() < s[j].GetName() }

type ProcessSlice []*info.Process

func (s ProcessSlice) Len() int           { return len(s) }
func (s ProcessSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ProcessSlice) Less(i, j int) bool { return s[i].GetName() < s[j].GetName() }

type ShellSlice []*info.Shell

func (s ShellSlice) Len() int           { return len(s) }
func (s ShellSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ShellSlice) Less(i, j int) bool { return s[i].GetName() < s[j].GetName() }

//检查每个服务状态是否变化，如果变化则产生retry
func check_service_status(host string, now, old []*info.Service) {
	if len(now) != len(old) {
		return
	}
	sort.Sort(ServiceSlice(now))
	sort.Sort(ServiceSlice(old))
	for i := 0; i < len(now); i++ {
		if now[i].Status == old[i].Status {
			continue
		}
		var retry = Retry{fmt.Sprintf("%s/service?q=%s&name=%s",
			host, now[i].Net, now[i].Name), "service", now[i].Ok()}
		notify.Retry <- retry
	}
}

//检查每个进程状态是否变化，如果变化则产生retry
func check_process_status(host string, now, old []*info.Process) {
	if len(now) != len(old) {
		return
	}
	sort.Sort(ProcessSlice(now))
	sort.Sort(ProcessSlice(old))
	for i := 0; i < len(now); i++ {
		if now[i].Pid == old[i].Pid {
			continue
		}
		var retry = Retry{fmt.Sprintf("%s/process?q=%s",
			host, now[i].Name), "process", now[i].Ok()}
		notify.Retry <- retry
	}
}

//检查每个脚本状态是否变化，如果变化则产生retry
func check_shell_status(host string, now, old []*info.Shell) {
	if len(now) != len(old) {
		return
	}
	sort.Sort(ShellSlice(now))
	sort.Sort(ShellSlice(old))
	for i := 0; i < len(now); i++ {
		if now[i].Result == old[i].Result {
			continue
		}
		var retry = Retry{fmt.Sprintf("%s/custom?q=%s",
			host, now[i].Name), "shell", now[i].Ok()}
		notify.Retry <- retry
	}
}

//检查retry,确定是否需要产生警报,输出到out
func check_if_warn(in <-chan Retry, out chan<- string) {
	for retry := range in {
		switch retry.Class {
		case "status":
			_, err := http.Head(retry.URL)
			stat := false
			if err != nil {
				notify.err <- err
			} else {
				stat = true
			}
			if stat == retry.Status {
				notify.Warn <- fmt.Sprintf("Warn:[%s] URL: %s\nContent:\nstatus changed %t",
					retry.Class, retry.URL, stat)
			}
		case "service":
			b := readResponse(retry.URL)
			var service info.Service
			err := fromjson(b, &service)
			if err != nil {
				notify.err <- err
				continue
			}
			if retry.Status == service.Ok() {
				notify.Warn <- fmt.Sprintf("Warn:[%s],URL: %s\nContent:\n%s",
					retry.Class, retry.URL, service.String())
			}
		case "process":
			b := readResponse(retry.URL)
			var process info.Process
			err := fromjson(b, &process)
			if err != nil {
				notify.err <- err
				continue
			}
			if retry.Status == process.Ok() {
				notify.Warn <- fmt.Sprintf("Warn:[%s],URL: %s\nContent:\n%s",
					retry.Class, retry.URL, process.String())
			}
		case "shell":
			b := readResponse(retry.URL)
			var shell info.Shell
			err := fromjson(b, &shell)
			if err != nil {
				notify.err <- err
				continue
			}
			if retry.Status == shell.Ok() {
				notify.Warn <- fmt.Sprintf("Warn:[%s],URL: %s\nContent:\n%s",
					retry.Class, retry.URL, shell.String())
			}
		}
	}
}
