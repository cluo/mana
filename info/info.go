package info

import (
	"mana/cfg"
	"runtime"
)

var (
	cf     = cfg.Parse()
	numcpu = runtime.NumCPU()
)
