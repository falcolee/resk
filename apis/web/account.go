package web

import (
	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
	"moyutec.top/resk/infra/base"
	"moyutec.top/resk/services"
)

func init() {
	base.RegisterApi(&AccountApi{})
}

type AccountApi struct {
	service services.AccountService
}

func (a *AccountApi) Init() {
	a.service = services.GetAccountService()
	groupRouter := base.Iris().Party("/v1/account")
	groupRouter.Post("/create", a.createHandler)
}

//账户创建的接口: /v1/account/create
//POST body json
/*
{
	"UserId": "w123456",
	"Username": "测试用户1",
	"AccountName": "测试账户1",
	"AccountType": 0,
	"CurrencyCode": "CNY",
	"Amount": "100.11"
}

{
    "code": 1000,
    "message": "",
    "data": {
        "AccountNo": "1K1hrG0sQw7lDuF6KOQbMBe2o3n",
        "AccountName": "测试账户1",
        "AccountType": 0,
        "CurrencyCode": "CNY",
        "UserId": "w123456",
        "Username": "测试用户1",
        "Balance": "100.11",
        "Status": 1,
        "CreatedAt": "2019-04-18T13:26:51.895+08:00",
        "UpdatedAt": "2019-04-18T13:26:51.895+08:00"
    }
}
*/
func (a *AccountApi) createHandler(ctx iris.Context) {
	//获取请求参数，
	account := services.AccountCreatedDTO{}
	err := ctx.ReadJSON(&account)
	r := base.Res{
		Code: base.ResCodeOk,
	}
	if err != nil {
		r.Code = base.ResCodeRequestParamsError
		r.Message = err.Error()
		ctx.JSON(&r)
		logrus.Error(err)
		return
	}
	//执行创建账户的代码
	dto, err := a.service.CreateAccount(account)
	if err != nil {
		r.Code = base.ResCodeInnerServerError
		r.Message = err.Error()
		logrus.Error(err)
	}
	r.Data = dto
	ctx.JSON(&r)
}
