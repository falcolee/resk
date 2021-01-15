package infra

import "github.com/tietang/props/v3/kvs"

//资源启动器上下文，
// 用来在服务资源初始化、安装、启动和停止的生命周期中变量和对象的传递
type StarterContext map[string]interface{}

func (s StarterContext) SetProps(conf kvs.ConfigSource) {
	s["_conf"] = conf
}

func (s StarterContext) Props() kvs.ConfigSource {
	p := s["_conf"]
	if p == nil {
		panic("配置文件未被初始化")
	}
	return p.(kvs.ConfigSource)
}

type Starter interface {
	Init(ctx StarterContext)
	Setup(ctx StarterContext)
	Start(ctx StarterContext)
	StartBlocking() bool
}

//默认的空实现,方便资源启动器的实现
type BaseStarter struct {
}

type starterRegister struct {
	nonBlockingStarters []Starter
	blockingStarters    []Starter
}

func (r *starterRegister) AllStarters() []Starter {
	starters := make([]Starter, 0)
	starters = append(starters, r.nonBlockingStarters...)
	starters = append(starters, r.blockingStarters...)
	return starters
}

func (r *starterRegister) Register(starter Starter) {
	if starter.StartBlocking() {
		r.blockingStarters = append(r.blockingStarters, starter)
	} else {
		r.nonBlockingStarters = append(r.nonBlockingStarters, starter)
	}
}

var StarterRegister *starterRegister = &starterRegister{}

func Register(starter Starter) {
	StarterRegister.Register(starter)
}

func GetStarters() []Starter {
	return StarterRegister.AllStarters()
}

func (s *BaseStarter) Init(ctx StarterContext)  {}
func (s *BaseStarter) Setup(ctx StarterContext) {}
func (s *BaseStarter) Start(ctx StarterContext) {}
func (s *BaseStarter) StartBlocking() bool      { return false }
