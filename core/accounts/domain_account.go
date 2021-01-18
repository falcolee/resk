package accounts

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/segmentio/ksuid"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"moyutec.top/resk/infra/base"
	"moyutec.top/resk/services"
)

type accountDomain struct {
	account    Account
	accountLog AccountLog
}

func NewAccountDomain() *accountDomain {
	return new(accountDomain)
}

func (domain *accountDomain) createAccountLogNo() {
	domain.accountLog.LogNo = ksuid.New().Next().String()
}

func (domain *accountDomain) createAccountNo() {
	domain.account.AccountNo = ksuid.New().Next().String()
}

//创建流水的记录
func (domain *accountDomain) createAccountLog() {
	//通过account来创建流水，创建账户逻辑在前
	domain.accountLog = AccountLog{}
	domain.createAccountLogNo()
	domain.accountLog.TradeNo = domain.accountLog.LogNo
	//流水中的交易主体信息
	domain.accountLog.AccountNo = domain.account.AccountNo
	domain.accountLog.UserId = domain.account.UserId
	domain.accountLog.Username = domain.account.Username
	//交易对象信息
	domain.accountLog.TargetAccountNo = domain.account.AccountNo
	domain.accountLog.TargetUserId = domain.account.UserId
	domain.accountLog.TargetUsername = domain.account.Username

	//交易金额
	domain.accountLog.Amount = domain.account.Balance
	domain.accountLog.Balance = domain.account.Balance
	//交易变化属性
	domain.accountLog.Decs = "账户创建"
	domain.accountLog.ChangeType = services.AccountCreated
	domain.accountLog.ChangeFlag = services.FlagAccountCreated
}

func (domain *accountDomain) Create(dto services.AccountDTO) (*services.AccountDTO, error) {
	domain.account = Account{}
	domain.account.FromDTO(&dto)
	domain.createAccountNo()
	domain.createAccountLog()
	accountDao := AccountDao{}
	accountLogDao := AccountLogDao{}
	var rdto *services.AccountDTO
	err := base.Tx(func(db *gorm.DB) error {
		accountDao.runner = db
		accountLogDao.runner = db
		id, err := accountDao.Insert(&domain.account)
		if err != nil {
			return err
		}
		if id <= 0 {
			return errors.New("创建账户流水失败")
		}
		domain.account = *accountDao.GetOne(domain.account.AccountNo)
		return nil
	})
	rdto = domain.account.ToDTO()
	return rdto, err
}

func (a *accountDomain) Transfer(dto services.AccountTransferDTO) (status services.TransferStatus, err error) {
	err = base.Tx(func(db *gorm.DB) error {
		status, err = a.TransferWithContext(dto)
		return err
	})
	return status, err
}

func (a *accountDomain) TransferWithContext(dto services.AccountTransferDTO) (status services.TransferStatus, err error) {
	amount := dto.Amount
	if dto.ChangeFlag == services.FlagTransferOut {
		amount = amount.Mul(decimal.NewFromFloat(-1))
	}

	a.accountLog = AccountLog{}
	a.accountLog.FromTransferDTO(&dto)
	a.createAccountLogNo()

	err = base.Tx(func(db *gorm.DB) error {
		accountDao := AccountDao{runner: db}
		accountLogDao := AccountLogDao{runner: db}
		rows, err := accountDao.UpdateBalance(dto.TradeBody.AccountNo, amount)
		if err != nil {
			status = services.TransferStatusFailure
		}
		if rows <= 0 && dto.ChangeFlag == services.FlagTransferOut {
			status = services.TransferStatusFundsNotAllowed
			return errors.New("余额不足")
		}
		account := accountDao.GetOne(dto.TradeBody.AccountNo)
		if account == nil {
			return errors.New("红包账户不存在")
		}
		a.account = *account
		a.accountLog.Balance = a.account.Balance
		id, err := accountLogDao.Insert(&a.accountLog)
		if err != nil || id <= 0 {
			status = services.TransferStatusFailure
			return errors.New("账户流水创建失败")
		}
		return nil
	})
	if err != nil {
		logrus.Error(err)
	} else {
		status = services.TransferStatusSuccess
	}
	return status, err
}

func (a *accountDomain) GetAccount(accountNo string) *services.AccountDTO {
	accountDao := AccountDao{runner: base.ORM()}
	account := accountDao.GetOne(accountNo)
	if account == nil {
		return nil
	}
	return account.ToDTO()
}

func (a *accountDomain) GetEnvelopeAccountByUserId(userId string) *services.AccountDTO {
	accountDao := AccountDao{runner: base.ORM()}
	account := accountDao.GetByUserId(userId, int(services.EnvelopeAccountType))
	if account == nil {
		return nil
	}
	return account.ToDTO()
}

func (a *accountDomain) GetAccountByUserIdAndType(userId string, accountType services.AccountType) *services.AccountDTO {
	accountDao := AccountDao{runner: base.ORM()}
	account := accountDao.GetByUserId(userId, int(accountType))
	if account == nil {
		return nil
	}
	return account.ToDTO()
}

func (a *accountDomain) GetAccountLog(logNo string) *services.AccountLogDTO {
	accountLogDao := AccountLogDao{runner: base.ORM()}
	accountLog := accountLogDao.GetOne(logNo)
	if accountLog == nil {
		return nil
	}
	return accountLog.ToDTO()
}

func (a *accountDomain) GetAccountLogByTradeNo(tradeNo string) *services.AccountLogDTO {
	accountLogDao := AccountLogDao{runner: base.ORM()}
	accountLog := accountLogDao.GetByTradeNo(tradeNo)
	if accountLog == nil {
		return nil
	}
	return accountLog.ToDTO()
}
