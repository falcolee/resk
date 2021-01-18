package base

import (
	"github.com/sirupsen/logrus"
	"moyutec.top/resk/infra"
	"net"
	"net/rpc"
	"reflect"
)

var rpcServer *rpc.Server

func RpcServer() *rpc.Server {
	return rpcServer
}

func RpcRegister(ri interface{}) {
	typ := reflect.TypeOf(ri)
	logrus.Infof("gorpc register:%s", typ.String())
	RpcServer().Register(ri)
}

type GoRPCStarter struct {
	infra.BaseStarter
}

func (s *GoRPCStarter) Start(ctx infra.StarterContext) {
	server := rpc.NewServer()

	port := ctx.Props().GetDefault("app.rpc.port", "8082")
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		logrus.Panic(err)
	}
	logrus.Info("tcp rpc port listened:", port)
	go server.Accept(listener)
	rpcServer = server
}
