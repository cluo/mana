/*
mana_agent 是一个简单的http服务器.

Options:
    -h=false: help infomation
    -i=false: enable indent to format the data
    -log="log/mana.log": log
    -p="": the password to encrypt data

主要的url:
    /system 系统的基本系统，负载以及流量温度等；
    /service 针对监听tcp和udp端口的服务检测，如/service?q=tcp|udp；
    /process 检查进程pid，用于检测特定的服务；
    /custom 可以编辑简单的脚本打印需要的信息；
    /stat 主要用来检查agent在线；
    /top 查找top以及内存使用最多的10个进程；
    /error_page ...
*/
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"mana/info"
	"net/http"
)

var agent *info.Agent
var cnf *config

//tcp/udp 地址:端口
type Address struct {
	Addr, Port string
}

//config 包含./etc/(tcp、udp、process、shell), 格式为json文本文件
type config struct {
	tcp, udp       map[string]Address
	process, shell []string
}

var (
	help = flag.Bool("h", false, "help infomation")
	//数据是否缩进为可读样式
	indent   = flag.Bool("i", false, "enable indent to format the data")
	iscrypto = flag.String("p", "", "the password to encrypt data")
	//info.Agent.Log 指定的日志路径，默认当前目录下的log/mana.log
	logfile = flag.String("log", "log/mana.log", "log path")
)

//读取配置文件需要检查的服务，进程，脚本等
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

func aesKey(s string) []byte {
	hash := md5.New()
	hash.Write([]byte(s))
	return hash.Sum(nil)
}

//golang json 结尾只能是' ','\n','\r','\t';
func aesEncrypt(block cipher.Block, src []byte) []byte {
	var dst = make([]byte, 16)
	var fill = []byte("                ")
	var src_len = len(src)
	if src_len%16 != 0 {
		src = append(src, fill[src_len%16:]...)
	}
	var enc []byte
	for i := 0; i < src_len; i += 16 {
		block.Encrypt(dst, src[i:i+16])
		enc = append(enc, dst...)
	}
	return enc
}

var block cipher.Block

//格式化数据为json
func tojson(v interface{}) ([]byte, error) {
	var bs []byte
	var err error
	if *indent {
		bs, err = json.MarshalIndent(v, "", "  ")
	} else {
		bs, err = json.Marshal(v)
	}
	//数据是否需要加密
	if block != nil {
		return aesEncrypt(block, bs), err
	}
	return bs, err
}

func system(w http.ResponseWriter, r *http.Request) {
	bts, err := tojson(agent.System())
	if err != nil {
		w.WriteHeader(503)
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

//对于/service?tcp|udp等，不设置name或者name为"all", 检查etc/tcp或者etc/udp指定的所有服务,
//如果name指定值，检查config.tcp或者config.udp，并返回json数据
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

//process etc/process config.process
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
				var hEAD bool
				p, err := agent.Process(name)
				if err != nil {
					w.WriteHeader(503)
					hEAD = true
					fmt.Fprintf(w, "%s", err)
				}
				bs, err := tojson(p)
				if err != nil {
					//是否以返回状态码
					if !hEAD {
						w.WriteHeader(503)
					}
					fmt.Fprintf(w, "%s", err)
					return
				}
				w.Header().Set("Mana-Status", "OK")
				fmt.Fprintf(w, "%s", bs)
				return
			}
		}
		w.WriteHeader(400)
		fmt.Fprintf(w, "No monitoring process: name %s", name)
	}
}

//custom etc/shell config.shell
func custom(w http.ResponseWriter, r *http.Request) {
	if len(cnf.shell) == 0 {
		w.WriteHeader(400)
		fmt.Fprint(w, "No custom output")
		return
	}
	query := r.URL.Query()
	//url example:
	//  /custom?q="myip" ...
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
				var hEAD bool
				sh, err := agent.Shell(name)
				if err != nil {
					w.WriteHeader(503)
					hEAD = true
					fmt.Fprintf(w, "%s", err)
				}
				bs, err := tojson(sh)
				if err != nil {
					if !hEAD {
						w.WriteHeader(503)
					}
					fmt.Fprintf(w, "%s", err)
					return
				}
				w.Header().Set("Mana-Status", "OK")
				fmt.Fprintf(w, "%s", bs)
				return
			}
		}
		w.WriteHeader(400)
		fmt.Fprintf(w, "No custom output: name %s", name)
	}
}

//主机运行时间以及loadavg
func stat(w http.ResponseWriter, r *http.Request) {
	h, err := agent.Hostname()
	if err != nil {
		w.WriteHeader(407)
		fmt.Fprintf(w, "%s", err)
		return
	}
	bootime := h.Boot.Format(info.Timestr)
	loadavg, err := agent.Loadavg()
	if err != nil {
		w.WriteHeader(407)
		fmt.Fprintln(w, err)
		return
	}
	w.Header().Set("Mana-Status", "OK")
	var s = fmt.Sprintf(`主机名: %s
启动时间: %s
已运行: %s
---------------------------------------------
系统负载: %s`, h.Name, bootime, h.Uptime, loadavg)
	//数据是否需要加密
	if block != nil {
		s = string(aesEncrypt(block, []byte(s)))
	}
	fmt.Fprintf(w, "%s", s)
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
	w.Header().Set("Mana-Status", "OK")
	var s = fmt.Sprintf("cpu使用 %s\n内存使用 %s", top_cpu, top_mem)
	if block != nil {
		s = string(aesEncrypt(block, []byte(s)))
	}
	fmt.Fprintf(w, "%s", s)
}

//others 404
func root(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI != "/" {
		http.NotFound(w, r)
		return
	}
	var remote_address string
	for i, remote := 0, r.RemoteAddr; i < len(remote); i++ {
		if remote[i] == ':' {
			remote_address = remote[:i]
			break
		}
	}
	w.Header().Set("Mana-Status", "OK")
	fmt.Fprint(w, remote_address)
}

func init() {
	flag.Parse()
	if *iscrypto != "" && len(*iscrypto) > 5 {
		var err error
		block, err = aes.NewCipher(aesKey(*iscrypto))
		if err != nil {
			panic(err)
		}
	}
	cnf = new(config)
	agent = info.NewAgent(info.NewLogger(*logfile))
	readFile("./etc/tcp", &cnf.tcp)
	readFile("./etc/udp", &cnf.udp)
	readFile("./etc/process", &cnf.process)
	readFile("./etc/shell", &cnf.shell)
}

func main() {
	if *help {
		//打印帮助
		flag.PrintDefaults()
		return
	}

	if *iscrypto != "" && len(*iscrypto) < 6 {
		fmt.Println("密码长度最少为6位")
		return
	}
	//
	http.HandleFunc("/custom", custom)
	http.HandleFunc("/service", service)
	http.HandleFunc("/system", system)
	http.HandleFunc("/process", process)
	http.HandleFunc("/stat", stat)
	http.HandleFunc("/top", top10)
	http.HandleFunc("/", root)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
