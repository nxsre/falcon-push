package main

import (
	"encoding/json"
	"fmt"
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func InitLogger() {
	// 日志地址 "out.log" 自定义
	lp := cfg.APP.LogFile
	// 日志级别 DEBUG,ERROR, INFO
	lv := cfg.APP.LogLevel
	// 是否输出到终端
	isStdout := true
	if cfg.APP.LogStdout != true {
		isStdout = false
	}
	initLogger(lp, lv, isStdout)
	log.SetFlags(log.Lmicroseconds | log.Lshortfile | log.LstdFlags)
}

func initLogger(lp string, lv string, isStdout bool) {
	var js string
	if isStdout {
		js = fmt.Sprintf(`{
      "level": "%s",
      "encoding": "json",
      "outputPaths": ["stdout"],
      "errorOutputPaths": ["stdout"]
      }`, lv)
	} else {
		js = fmt.Sprintf(`{
      "level": "%s",
      "encoding": "json",
      "outputPaths": ["%s"],
      "errorOutputPaths": ["%s"]
      }`, lv, lp, lp)
	}

	var zapcfg zap.Config
	if err := json.Unmarshal([]byte(js), &zapcfg); err != nil {
		panic(err)
	}

	zapcfg.EncoderConfig = zap.NewProductionEncoderConfig()
	zapcfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	var err error
	logger, err = zapcfg.Build()
	if err != nil {
		log.Fatal("init logger error: ", err)
	}
}
