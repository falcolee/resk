package base

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	recover2 "github.com/kataras/iris/v12/middleware/recover"
	"github.com/sirupsen/logrus"
	"moyutec.top/resk/infra"
	"time"
)

var irisApplication *iris.Application

func Iris() *iris.Application {
	return irisApplication
}

type IrisServiceStart struct {
	infra.BaseStarter
}

func (i *IrisServiceStart) Init(ctx infra.StarterContext) {
	// 创建iris application实例
	irisApplication = initIris()
	// 日志组件配置和扩展
	irisLogger := irisApplication.Logger()
	irisLogger.Install(logrus.StandardLogger())
}

func (i *IrisServiceStart) Start(ctx infra.StarterContext) {
	routes := Iris().GetRoutes()
	for _, route := range routes {
		logrus.Info(route.Trace())
	}
	//启动
	Iris().Run(iris.Addr(":8082"))
}

func (i *IrisServiceStart) StartBlocking() bool {
	return true
}

func initIris() *iris.Application {
	app := iris.New()
	app.Use(recover2.New())
	cfg := logger.Config{
		Status: true,
		IP:     true,
		Method: true,
		Path:   true,
		Query:  true,
		LogFunc: func(endTime time.Time, latency time.Duration, status, ip, method, path string, message interface{}, headerMessage interface{}) {
			app.Logger().Infof("| %s | %s | %s | %s | %s | %s | %+v | %+v",
				endTime.Format("2006-01-02.15:04:05.000000"),
				latency.String(), status, ip, method, path, headerMessage, message,
			)
		},
	}
	app.Use(logger.New(cfg))
	return app
}
