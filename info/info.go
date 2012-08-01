/*
   简单的获取cpu利用率，系统内存使用情况，磁盘读写情况.
   查看cpu温度命令sensors输出内容，判断是否高于（70）等.
   获取硬盘温度.
   使用package net检查tcp/udp等服务是否在线.
   使用shell脚本检查特定进程。
*/
package info

import (
	"errors"
	"log"
	"os"
)

func NewLogger(name string) *log.Logger {
	file, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("opening log:", err)
	}
	return log.New(file, "", log.LstdFlags)
}

var elog = log.New(os.Stderr, "", log.LstdFlags)

type Server struct {
	Hostname *Hostname
	Load     *Load
	Temp     *Temp
}

type Agent struct{}

func (a *Agent) System() (*Server, error) {
	sys := new(Server)
	host, err := GetHostname()
	if err != nil {
		return nil, errors.New("func GetHostname() failed")
	}
	sys.Hostname = host
	load, err := GetLoad()
	if err != nil {
		return nil, errors.New("func GetLoad() failed")
	}
	sys.Load = load
	temp, err := GetTemp()
	if err != nil {
		return nil, errors.New("func GetTemp() failed")
	}
	sys.Temp = temp
	return sys, nil
}
