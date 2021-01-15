package resk

import (
	"moyutec.top/resk/infra"
	"moyutec.top/resk/infra/base"
)

func init() {
	infra.Register(&base.PropsStarter{})
	infra.Register(&base.ValidatorStarter{})
	infra.Register(&base.GormStarter{})
	infra.Register(&base.IrisServiceStart{})
}
