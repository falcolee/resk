package accounts

import (
	"github.com/shopspring/decimal"
	"moyutec.top/resk/services"
	"time"
)

type AccountLog struct {
	Id              int64               `gorm:"column:id;primaryKey"`     //
	LogNo           string              `gorm:"column:log_no;unique"`     //流水编号 全局不重复字符或数字，唯一性标识
	TradeNo         string              `gorm:"column:trade_no"`          //交易单号 全局不重复字符或数字，唯一性标识
	AccountNo       string              `gorm:"column:account_no"`        //账户编号 账户ID
	UserId          string              `gorm:"column:user_id"`           //用户编号
	Username        string              `gorm:"column:username"`          //用户名称
	TargetAccountNo string              `gorm:"column:target_account_no"` //账户编号 账户ID
	TargetUserId    string              `gorm:"column:target_user_id"`    //目标用户编号
	TargetUsername  string              `gorm:"column:target_username"`   //目标用户名称
	Amount          decimal.Decimal     `gorm:"column:amount"`            //交易金额,该交易涉及的金额
	Balance         decimal.Decimal     `gorm:"column:balance"`           //交易后余额,该交易后的余额
	ChangeType      services.ChangeType `gorm:"column:change_type"`       //流水交易类型，0 创建账户，>0 为收入类型，<0 为支出类型，自定义
	ChangeFlag      services.ChangeFlag `gorm:"column:change_flag"`       //交易变化标识：-1 出账 1为进账，枚举
	Status          int                 `gorm:"column:status"`            //交易状态：
	Decs            string              `gorm:"column:decs"`              //交易描述
	CreatedAt       time.Time           `gorm:"column:created_at"`        //创建时间
}

func (AccountLog) TableName() string {
	return "account_log"
}

func (po *AccountLog) FromTransferDTO(dto *services.AccountTransferDTO) {
	po.TradeNo = dto.TradeNo
	po.AccountNo = dto.TradeBody.AccountNo
	po.TargetAccountNo = dto.TradeTarget.AccountNo
	po.UserId = dto.TradeBody.UserId
	po.Username = dto.TradeBody.Username
	po.TargetUserId = dto.TradeTarget.UserId
	po.TargetUsername = dto.TradeTarget.Username
	po.Amount = dto.Amount
	po.ChangeType = dto.ChangeType
	po.ChangeFlag = dto.ChangeFlag
	po.Decs = dto.Decs
}

func (po *AccountLog) ToDTO() *services.AccountLogDTO {
	dto := &services.AccountLogDTO{

		TradeNo:         po.TradeNo,
		LogNo:           po.LogNo,
		AccountNo:       po.AccountNo,
		TargetAccountNo: po.TargetAccountNo,
		UserId:          po.UserId,
		Username:        po.Username,
		TargetUserId:    po.TargetUserId,
		TargetUsername:  po.TargetUsername,
		Amount:          po.Amount,
		Balance:         po.Balance,
		ChangeType:      po.ChangeType,
		ChangeFlag:      po.ChangeFlag,
		Status:          po.Status,
		Decs:            po.Decs,
		CreatedAt:       po.CreatedAt,
	}
	return dto
}
