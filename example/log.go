package example

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	formatter := &log.TextFormatter{}
	formatter.FullTimestamp = true
	log.SetFormatter(formatter)
	level := os.Getenv("log.debug")
	if level == "true" {
		log.SetLevel(log.DebugLevel)
	}
	formatter.ForceColors = true
	formatter.DisableColors = false
	log.Debug("debug 测试")
	log.Info("测试")

}
