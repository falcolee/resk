package accounts

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type AccountLogDao struct {
	runner *gorm.DB
}

func (dao *AccountLogDao) GetOne(logNo string) *AccountLog {
	a := &AccountLog{LogNo: logNo}
	result := dao.runner.Where(a).First(a)
	if result.RecordNotFound() {
		return nil
	} else if result.Error != nil {
		logrus.Error(result.Error)
		return nil
	}
	return a
}

func (dao *AccountLogDao) GetByTradeNo(tradeNo string) *AccountLog {
	a := &AccountLog{TradeNo: tradeNo}
	result := dao.runner.Where(a).First(a)
	if result.RecordNotFound() {
		return nil
	} else if result.Error != nil {
		logrus.Error(result.Error)
		return nil
	}
	return a
}

func (dao *AccountLogDao) Insert(accountLog *AccountLog) (id int64, err error) {
	result := dao.runner.Create(accountLog)
	err = result.Error
	return accountLog.Id, err
}
