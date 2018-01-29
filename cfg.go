package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Cfg struct {
	APP      APP      `yaml:"app"`
	DingTalk DingTalk `yaml:"dingtalk"`
	DingEnt  DingEnt  `yaml:"dingent"`
	SMTP     SMTP     `yaml:"smtp"`
	SMS      SMS      `yaml:"sms"`
}

type APP struct {
	Listen    string `yaml:"listen"`
	LogLevel  string `yaml:"loglevel"`
	LogFile   string `yaml:"logfile"`
	LogStdout bool   `yaml:"logstdout"`
	ClientId  string `yaml:"client_id"`
}

type SMS struct {
	URL string `yaml:"url"`
}

type DingTalk struct {
	URL string `yaml:"url"`
}

type DingEnt struct {
	URL string `yaml:"url"`
}

type SMTP struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
}

var cfg *Cfg

func init() {
	if cfgb, err := ioutil.ReadFile("config.yml"); err != nil {
		log.Fatalln(err)
	} else {
		if err := yaml.Unmarshal(cfgb, &cfg); err != nil {
			log.Fatalln(err)
		}
	}
	InitLogger()
}
