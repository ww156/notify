/*
# @Author ww
# @Time 2019/10/28 10:01
# @File log.go
*/
package log

import (
	"encoding/json"
	"go.uber.org/zap"
)

var (
	Logger *zap.Logger
)

func init() {
	rawJSON := []byte(`{
	  "level": "debug",
	  "encoding": "json",
	  "outputPaths": ["stdout", "/tmp/logs"],
	  "errorOutputPaths": ["stderr"],
	}`)
	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	Logger = logger
}
