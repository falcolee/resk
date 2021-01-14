package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func main() {
	app := iris.Default()
	app.Get("/hello", func(context context.Context) {
		context.WriteString("hello,world")
	})
	app.Get("/error", func(ctx context.Context) {
		panic("出错了")
	})
	v1 := app.Party("/v1")
	v1.Use(func(ctx context.Context) {
		logrus.Info("自定义中间件")
		ctx.Next()
	})
	v1.Get("/users/{id:uint64}", func(i context.Context) {
		val := i.Params().GetUint64Default("id", 0)
		fmt.Println(val)
		i.WriteString(strconv.Itoa(int(val)))
	})
	app.OnAnyErrorCode(func(ctx context.Context) {
		ctx.WriteString("出错了啊")
	})
	app.OnErrorCode(http.StatusNotFound, func(ctx context.Context) {
		ctx.WriteString("访问的路径不存在")
	})
	app.Run(iris.Addr(":8082"))
}
