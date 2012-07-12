package cfg

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Base      string
	System    bool
	Log       string
	Services  []*Service
	Processes []*Process
}

func Parse() *Config {
	var c = new(Config)

	bs, err := ioutil.ReadFile("../etc/config.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(bs, c)
	if err != nil {
		log.Fatal(err)
	}
	return c
}

type Log string

func (c *Config) Openlog() *log.Logger {
	lf, err := os.OpenFile(c.Log, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	return log.New(lf, "", log.LstdFlags)
}
