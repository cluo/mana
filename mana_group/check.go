package main

import (
	"fmt"
	"mana/info"
	"net/http"
	"regexp"
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

//检查系统负载
func check_sys_load(host string, now, old *info.Load) {
	if old == nil {
		return
	}
	var warn_load string
	var overload = 5.0
	//loadavg
	la5_now, la5_old := now.Loadavg.Load5(), old.Loadavg.Load5()
	if la5_now > overload && la5_old < overload {
		warn_load += fmt.Sprintf("loadavg(5) too high\n%s\n",
			now.Loadavg.String())
	} else if la5_now < overload && la5_old > overload {
		warn_load += fmt.Sprintf("loadavg(5) is fine\n%s\n",
			now.Loadavg.String())
	}
	//cpu
	for i := 0; i < len(old.Cpu); i++ {
		if now.Cpu[i].Idle < 5.0 && old.Cpu[i].Idle > 5.0 {
			warn_load += fmt.Sprintf("Cpu(s) utilization too high\n%s\n",
				now.Cpu[i].String())
		} else if now.Cpu[i].Idle > 5.0 && old.Cpu[i].Idle < 5.0 {
			warn_load += fmt.Sprintf("Cpu(s) utilization is fine\n%s\n",
				now.Cpu[i].String())
		}
	}
	//memory
	var less float64 = 104857600
	nr, or := now.Free.Real(), old.Free.Real()
	if nr < less && or > less {
		warn_load += fmt.Sprintf("memory free less than 100M\n%s\n",
			now.Free.Format().String())
	} else if nr > less && or < less {
		warn_load += fmt.Sprintf("memory free is fine\n%s\n",
			now.Free.Format().String())
	}
	//io
	if warn_load != "" && now.Cpu[0].Wa > 50 {
		warn_load += string(now.IO)
		warn_load += "\n"
	}

	if warn_load != "" {
		notify.Warn <- fmt.Sprintf("Warn:[load] Host: %s\n%s",
			host, warn_load)
	}
}

//正则截取传感器温度
var ReSensors = regexp.MustCompile(`.+?(\d+\.\d)°C\s+`)

//检查硬盘和cpu温度, 65|C
func check_sys_temp(host string, now, old *info.Temp) {
	var high = "65"
	var warn_temp string
	//hddtemp
	now_hdd, old_hdd := now.Disks, old.Disks
	for i := 0; i < len(old.Disks); i++ {
		if now_hdd[i].Temp != "UNK" && old_hdd[i].Temp != "UNK" {
			if now_hdd[i].Temp > high && old_hdd[i].Temp < high {
				warn_temp += fmt.Sprintf("hddtemp high than %s\n%s\n",
					high, now_hdd[i])
			} else if now_hdd[i].Temp < high && old_hdd[i].Temp > high {
				warn_temp += fmt.Sprintf("hddtemp is fine\n%s\n", now_hdd[i])
			}
		}
	}
	//sensors
	now_sensors := ReSensors.FindAllStringSubmatch(string(now.Sensors), -1)
	old_sensors := ReSensors.FindAllStringSubmatch(string(old.Sensors), -1)
	for i := 0; i < len(old_sensors); i++ {
		now_s, old_s := now_sensors[i], old_sensors[i]
		if now_s[1] > high && old_s[1] < high {
			warn_temp += fmt.Sprintf("cpu sensors high than %s\n%s\n",
				high, now_s[0])
		} else if now_s[1] < high && old_s[1] > high {
			warn_temp += fmt.Sprintf("cpu sensors is fine\n%s\n", now_s[0])
		}
	}
	if warn_temp != "" {
		notify.Warn <- fmt.Sprintf("Warn:[temp] Host: %s\n%s",
			host, warn_temp)
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
				notify.Warn <- fmt.Sprintf("Warn:[%s] URL: %s\nContent:\n%s",
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
				notify.Warn <- fmt.Sprintf("Warn:[%s] URL: %s\nContent:\n%s",
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
				notify.Warn <- fmt.Sprintf("Warn:[%s] URL: %s\nContent:\n%s",
					retry.Class, retry.URL, shell.String())
			}
		}
	}
}
