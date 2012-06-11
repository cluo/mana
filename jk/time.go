package jk

import (
    "time"
)

type Timer interface {
    Stat() []byte
    String() string
    Interval() time.Duration
    Last(time.Time)
}

func SetDuration(t Timer, out chan<- Timer) {
    var td = time.Duration(t.Interval())*time.Second
    var ch = time.Tick(td)
    for n :=range ch {
        t.Last(n)
        out <- t
    }
}
