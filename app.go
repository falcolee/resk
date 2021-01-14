package resk

import (
	"moyutec.top/resk/infra"
	"moyutec.top/resk/infra/base"
)

func init() {
	infra.Register(&base.PropsStarter{})
	infra.Register(&base.IrisServiceStart{})
}
