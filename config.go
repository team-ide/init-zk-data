package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type InitConfig struct {
	Server string
	Data   map[interface{}]interface{}
}

var (
	Config *InitConfig
)

func init() {
	filePath := "./config.yml"
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	Config = &InitConfig{}
	err = yaml.Unmarshal(bytes, Config)
	if err != nil {
		panic(err)
	}
}
