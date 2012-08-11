package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"mana/info"
	"net/http"
	_ "net/http/pprof"
)

var agent *info.Agent
var cnf *config

type Address struct {
	Addr, Port string
}

type config struct {
	tcp, udp       map[string]Address
	process, shell []string
}

var indent = flag.Bool("i", false, "json strings indent")
var logfile = flag.String("log", "./log/mana.log", "log path")

func readFile(path string, v interface{}) {
	fs, err := ioutil.ReadFile(path)
	if err != nil {
		agent.Log.Println(path, "not exists, ignored")
		return
	} else {
		err = json.Unmarshal(fs, v)
		if err != nil {
			panic(err)
		}
	}
}

func tojson(v interface{}) ([]byte, error) {
	if *indent {
		return json.MarshalIndent(v, "", "  ")
	}
	return json.Marshal(v)
}

func system(w http.ResponseWriter, r *http.Request) {
	bts, err := tojson(agent.System())
	if err != nil {
		fmt.Fprint(w, "json marshal error", err)
		return
	}
	w.Header().Set("Mana-Status", "OK")
	fmt.Fprintf(w, "%s", bts)
}

func tcp(name, addr, port string) *info.Service {
	return agent.Tcp(name, addr, port)
}
func udp(name, addr, port string) *info.Service {
	return agent.Udp(name, addr, port)
}

func checkService(listen map[string]Address, fn func(string, string, string) *info.Service, name string) ([]byte, error) {
	if len(listen) > 0 {
		switch name {
		case "", "all":
			var sr []*info.Service
			for k, v := range listen {
				sr = append(sr, fn(k, v.Addr, v.Port))
			}
			return tojson(sr)
		default:
			if v, ok := listen[name]; ok {
				return tojson(fn(name, v.Addr, v.Port))
			}
			return nil, fmt.Errorf("No monitoring service: name %s", name)
		}
	}
	return nil, fmt.Errorf("No monitoring any services")
}

func service(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	var protocol, name = query.Get("q"), query.Get("name")
	switch protocol {
	case "tcp":
		bs, err := checkService(cnf.tcp, tcp, name)
		if err != nil {
			w.WriteHeader(400)
			fmt.Fprintf(w, "%s", err)
			return
		}
		w.Header().Set("Mana-Status", "OK")
		fmt.Fprintf(w, "%s", bs)
	case "udp":
		bs, err := checkService(cnf.udp, udp, name)
		if err != nil {
			w.WriteHeader(400)
			fmt.Fprintf(w, "%s", err)
			return
		}
		w.Header().Set("Mana-Status", "OK")
		fmt.Fprintf(w, "%s", bs)
	default:
		w.WriteHeader(400)
		fmt.Fprintf(w, "Invalid parameters: %s\nMust: /service?q=tcp|udp[&name=all|\"name\"]", r.URL)
	}
}

func process(w http.ResponseWriter, r *http.Request) {
	if len(cnf.process) == 0 {
		w.WriteHeader(400)
		fmt.Fprint(w, "No monitoring any processes")
		return
	}
	query := r.URL.Query()
	var name = query.Get("q")
	switch name {
	case "", "all":
		var procs []*info.Process
		for _, v := range cnf.process {
			p, _ := agent.Process(v)
			procs = append(procs, p)
		}
		bs, err := tojson(procs)
		if err != nil {
			w.WriteHeader(503)
			fmt.Fprintf(w, "%s", err)
			return
		}
		w.Header().Set("Mana-Status", "OK")
		fmt.Fprintf(w, "%s", bs)
	default:
		for _, v := range cnf.process {
			if name == v {
				p, _ := agent.Process(name)
				bs, err := tojson(p)
				if err != nil {
					w.WriteHeader(503)
					fmt.Fprintf(w, "%s", err)
					return
				}
				fmt.Fprintf(w, "%s", bs)
				return
			}
		}
		w.WriteHeader(400)
		fmt.Fprintf(w, "No monitoring process: name %s", name)
	}
}

func custom(w http.ResponseWriter, r *http.Request) {
	if len(cnf.shell) == 0 {
		w.WriteHeader(400)
		fmt.Fprint(w, "No custom output")
		return
	}
	query := r.URL.Query()
	var name = query.Get("q")
	switch name {
	case "", "all":
		var shells []*info.Shell
		for _, v := range cnf.shell {
			sh, _ := agent.Shell(v)
			shells = append(shells, sh)
		}
		bs, err := tojson(shells)
		if err != nil {
			w.WriteHeader(503)
			fmt.Fprintf(w, "%s", err)
			return
		}
		w.Header().Set("Mana-Status", "OK")
		fmt.Fprintf(w, "%s", bs)
	default:
		for _, v := range cnf.shell {
			if name == v {
				sh, _ := agent.Shell(name)
				bs, err := tojson(sh)
				if err != nil {
					w.WriteHeader(503)
					fmt.Fprintf(w, "%s", err)
					return
				}
				fmt.Fprintf(w, "%s", bs)
				return
			}
		}
		w.WriteHeader(400)
		fmt.Fprintf(w, "No custom output: name %s", name)
	}
}

func status(w http.ResponseWriter, r *http.Request) {
	h, err := agent.Hostname()
	if err != nil {
		w.WriteHeader(407)
		fmt.Fprintf(w, "%s", err)
		return
	}
	bootime := h.Boot.Format(info.Timestr)
	fmt.Fprintf(w, "主机名: %s\n启动时间: %s\n已运行: %s", h.Name, bootime, h.Uptime)
	fmt.Fprintln(w, "\n---------------------------------------------")
	loadavg, err := agent.Loadavg()
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	fmt.Fprintf(w, "系统负载: %s", loadavg)
	fmt.Fprintln(w, "\nstatus: OK")
}

func top10(w http.ResponseWriter, r *http.Request) {
	top_cpu, err := agent.Top("10", "cpu")
	var hEAD bool
	if err != nil {
		w.WriteHeader(503)
		hEAD = true
		fmt.Fprintf(w, "%s", err)
	}
	top_mem, err := agent.Top("10", "mem")
	if err != nil {
		if !hEAD {
			w.WriteHeader(503)
		}
		fmt.Fprintf(w, "%s", err)
	}
	fmt.Fprintf(w, "cpu使用 %s\n内存使用 %s", top_cpu, top_mem)
}

func init() {
	flag.Parse()
	cnf = new(config)
	agent = info.NewAgent(info.NewLogger(*logfile))
	readFile("./etc/tcp", &cnf.tcp)
	readFile("./etc/udp", &cnf.udp)
	readFile("./etc/process", &cnf.process)
	readFile("./etc/shell", &cnf.shell)
}

func main() {
	http.HandleFunc("/custom", custom)
	http.HandleFunc("/service", service)
	http.HandleFunc("/system", system)
	http.HandleFunc("/process", process)
	http.HandleFunc("/status", status)
	http.HandleFunc("/top", top10)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
