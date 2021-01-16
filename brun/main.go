package main

import (
	"github.com/tietang/props/v3/ini"
	"github.com/tietang/props/v3/kvs"
	_ "moyutec.top/resk"
	"moyutec.top/resk/infra"
	"moyutec.top/resk/infra/base"
)

func main() {
	file := kvs.CurrentFilePath("config.ini", 1)
	conf := ini.NewIniFileCompositeConfigSource(file)
	base.InitLog(conf)
	app := infra.New(conf)
	app.Start()
}
