package accounts

import (
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type AccountDao struct {
	runner *gorm.DB
}

func (dao *AccountDao) GetOne(accountNo string) *Account {
	a := &Account{AccountNo: accountNo}
	result := dao.runner.Where(a).First(a)
	if result.Error != nil {
		logrus.Error(result.Error)
		return nil
	}
	return a
}

func (dao *AccountDao) GetByUserId(userId string, accountType int) *Account {
	a := &Account{UserId: userId, AccountType: accountType}
	result := dao.runner.Where(a).First(a)
	if result.Error != nil {
		logrus.Error(result.Error)
		return nil
	}
	return a
}

func (dao *AccountDao) Insert(account *Account) (id int64, err error) {
	result := dao.runner.Create(account)
	err = result.Error
	return account.Id, err
}

//账户余额的更新
//amount 如果是负数，就是扣减；如果是正数，就是增加
func (dao *AccountDao) UpdateBalance(
	accountNo string,
	amount decimal.Decimal) (rows int64, err error) {
	rs := dao.runner.Model(Account{}).Where("account_no=?", accountNo).UpdateColumn("balance", gorm.Expr("balance + ?", amount.String()))
	if rs.Error != nil {
		return 0, rs.Error
	}
	return rs.RowsAffected, nil
}

func (dao *AccountDao) UpdateStatus(
	accountNo string,
	status int) (rows int64, err error) {
	rs := dao.runner.Model(Account{}).Where("account_no=?", accountNo).Update("status", status)
	if rs.Error != nil {
		return 0, rs.Error
	}
	return rs.RowsAffected, nil
}
