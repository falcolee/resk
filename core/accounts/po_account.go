package accounts

import (
	"github.com/shopspring/decimal"
	"time"
)

// 账户持久化对象
type Account struct {
	Id           int64           `gorm:"column:id;primaryKey"`     //账户ID
	AccountNo    string          `gorm:"column:account_no;unique"` //账户编号,账户唯一标识
	AccountName  string          `gorm:"column:account_name"`      //账户名称,用来说明账户的简短描述,账户对应的名称或者命名，比如xxx积分、xxx零钱
	AccountType  int             `gorm:"column:account_type"`      //账户类型，用来区分不同类型的账户：积分账户、会员卡账户、钱包账户、红包账户
	CurrencyCode string          `gorm:"column:currency_code"`     //货币类型编码：CNY人民币，EUR欧元，USD美元 。。。
	UserId       string          `gorm:"column:user_id"`           //用户编号, 账户所属用户
	Username     string          `gorm:"column:username"`          //用户名称
	Balance      decimal.Decimal `gorm:"column:balance"`           //账户可用余额
	Status       int             `gorm:"column:status"`            //账户状态，账户状态：0账户初始化，1启用，2停用
	CreatedAt    time.Time       `gorm:"column:created_at"`        //创建时间
	UpdatedAt    time.Time       `gorm:"column:updated_at"`        //更新时间
}

func (Account) TableName() string {
	return "account"
}
