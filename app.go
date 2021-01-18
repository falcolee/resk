package resk

import (
	_ "moyutec.top/resk/apis/web"
	"moyutec.top/resk/infra"
	"moyutec.top/resk/infra/base"
)

func init() {
	infra.Register(&base.PropsStarter{})
	infra.Register(&base.ValidatorStarter{})
	infra.Register(&base.GormStarter{})
	infra.Register(&base.IrisServiceStart{})
	infra.Register(&base.WebApiStarter{})
	infra.Register(&base.GoRPCStarter{})
}
