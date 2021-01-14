package base

import (
	"github.com/sirupsen/logrus"
	"moyutec.top/resk/infra"
)

type PropsStarter struct {
	infra.BaseStarter
}

func (p *PropsStarter) Init(ctx infra.StarterContext) {
	logrus.Info("初始化配置项")
}

func (p *PropsStarter) Start(ctx infra.StarterContext) {
}

func (p *PropsStarter) StartBlocking() bool {
	return false
}
