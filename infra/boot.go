package infra

import (
	"github.com/sirupsen/logrus"
	"github.com/tietang/props/v3/kvs"
)

type BootApplication struct {
	IsTest     bool
	starterCtx StarterContext
	conf       kvs.ConfigSource
}

func (b *BootApplication) Start() {
	b.init()
	b.setup()
	b.start()
}

func (b *BootApplication) init() {
	logrus.Info("Init starters")
	for _, starter := range GetStarters() {
		starter.Init(b.starterCtx)
	}
}

func (b *BootApplication) setup() {
	logrus.Info("setup starters")
	for _, starter := range GetStarters() {
		starter.Setup(b.starterCtx)
	}
}

func (b *BootApplication) start() {
	logrus.Info("starting starters...")
	for k, starter := range GetStarters() {
		if starter.StartBlocking() {
			if k+1 == len(GetStarters()) {
				starter.Start(b.starterCtx)
			} else {
				go starter.Start(b.starterCtx)
			}
		} else {
			starter.Start(b.starterCtx)
		}
	}
}

func New(conf kvs.ConfigSource) *BootApplication {
	e := &BootApplication{
		IsTest:     false,
		starterCtx: StarterContext{},
		conf:       conf,
	}
	e.starterCtx.SetProps(conf)
	return e
}
