package testx

import (
	"github.com/tietang/props/v3/ini"
	"github.com/tietang/props/v3/kvs"
	"moyutec.top/resk/infra"
	"moyutec.top/resk/infra/base"
)

func init() {
	file := kvs.CurrentFilePath("../../brun/config.ini", 1)
	conf := ini.NewIniFileCompositeConfigSource(file)
	base.InitLog(conf)

	infra.Register(&base.PropsStarter{})
	infra.Register(&base.GormStarter{})
	infra.Register(&base.ValidatorStarter{})

	app := infra.New(conf)
	app.Start()
}
