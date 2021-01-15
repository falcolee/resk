package main

import (
	"github.com/tietang/props/v3/ini"
	_ "moyutec.top/resk"
	"moyutec.top/resk/infra"
	"moyutec.top/resk/infra/base"
)

func main() {
	conf := ini.NewIniFileCompositeConfigSource("/Users/legendol/go/src/resk/brun/config.ini")
	base.InitLog(conf)
	app := infra.New(conf)
	app.Start()
}
