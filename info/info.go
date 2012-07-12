/* 使用mana/cfg中的config；
 * 简单的获取cpu利用率，系统内存使用情况，磁盘读写情况；
 * 查看cpu温度命令sensors输出内容，判断是否高于（70）等；
 * 获取硬盘温度；
 * 使用package net检查tcp/udp等服务是否在线；
 * 使用shell脚本检查特定进程。
 *
 */
package info

import (
	"mana/cfg"
	"runtime"
)
// mana.Config 读取
var (
	cf     = cfg.Parse()
	numcpu = runtime.NumCPU()
)
