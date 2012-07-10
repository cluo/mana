package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Basedir    string
    Log *Log
    Services []*Service
    Processes []*Process
}

func Parse() *Config {
    var config = new(Config)

	bs, err := ioutil.ReadFile("../etc/config")
	if err != nil {
        log.Fatal(err)
	}
	err = json.Unmarshal(bs, config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}
