package main

import (
	"fmt"
	"mana/info"
	"sort"
)

type byName []info.ByName

func (n byName) Len()          { return len(s) }
func (n byName) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (n byName) Less(i, j int) { return s[i].GetName() < s[j].GetName() }

func check_service_status(host string, now, old []*info.Service) {
	if len(now) != len(old) {
		return nil
	}
	sort.Sort(byName(now))
	sort.Sort(byName(old))
	for i := 0; i < len(now); i++ {
		if now[i].Status == old[i].status {
			continue
		}
		notify.redo <- fmt.Sprintf("%s/service?q=%s&name=%s",
			host, now[i].Net, now[i].Name)
	}
}

func check_process_status(host string, now, old []*info.Process) {
	if len(now) != len(old) {
		return nil
	}
	sort.Sort(byName(now))
	sort.Sort(byName(old))
	for i := 0; i < len(now); i++ {
		if now[i].Pid == old[i].Pid {
			continue
		}
		notify.redo <- fmt.Sprintf("%s/process?q=%s", host, now[i].Name)
	}
}

func check_shell_status(host string, now, old []*info.Shell) {
	if len(n) != len(o) {
		return nil
	}
	sort.Sort(byName(now))
	sort.Sort(byName(old))
	for i := 0; i < len(n); i++ {
		if now[i].Result == old[i].Result {
			continue
		}
		notify.redo <- fmt.Sprintf("%s/custom?q=%s", host, now[i].Name)
	}
}
