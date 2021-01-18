package base

import (
	"moyutec.top/resk/infra"
)

//初始化接口

type Initializer interface {
	//用于对象实例化后的初始化操作
	Init()
}

//初始化注册器
type InitializeRegister struct {
	Initializers []Initializer
}

//注册一个初始化对象
func (i *InitializeRegister) Register(ai Initializer) {
	i.Initializers = append(i.Initializers, ai)
}

var apiInitializerRegister = new(InitializeRegister)

func GetApiInitializers() []Initializer {
	return apiInitializerRegister.Initializers
}

func RegisterApi(ai Initializer) {
	apiInitializerRegister.Register(ai)
}

type WebApiStarter struct {
	infra.BaseStarter
}

func (a *WebApiStarter) Setup(ctx infra.StarterContext) {
	for _, initializer := range GetApiInitializers() {
		initializer.Init()
	}
}
