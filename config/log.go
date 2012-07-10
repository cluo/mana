package config

import (
    "log"
    "os"
)

type Log struct {
    Error string
    Info string
}

func (l *Log) Open() (e, i *log.Logger) {

    elog, err := os.OpenFile(l.Error,os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
    if err != nil {
        log.Fatal(err)
    }
    ilog, err := os.OpenFile(l.Info, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
    if err != nil {
        log.Fatal(err)
    }
    e = log.New(elog,"",log.LstdFlags)
    i = log.New(ilog,"",log.LstdFlags)
    return
}
