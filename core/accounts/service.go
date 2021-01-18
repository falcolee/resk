package accounts

import (
	"errors"
	"github.com/shopspring/decimal"
	"moyutec.top/resk/infra/base"
	"moyutec.top/resk/services"
	"sync"
)

var _ services.AccountService = new(accountService)
var once sync.Once

func init() {
	once.Do(func() {
		services.IAccountService = new(accountService)
	})
}

type accountService struct {
}

func (a *accountService) GetAccount(accountNo string) *services.AccountDTO {
	domain := accountDomain{}
	return domain.GetAccount(accountNo)
}

func (a *accountService) CreateAccount(dto services.AccountCreatedDTO) (*services.AccountDTO, error) {
	domain := accountDomain{}
	if err := base.ValidateStruct(&dto); err != nil {
		return nil, err
	}

	acc := domain.GetAccountByUserIdAndType(dto.UserId, services.AccountType(dto.AccountType))
	if acc != nil {
		return acc, errors.New("用户的该类型账户已经存在")
	}
	amount, err := decimal.NewFromString(dto.Amount)
	if err != nil {
		return nil, err
	}
	account := services.AccountDTO{
		UserId:       dto.UserId,
		Username:     dto.Username,
		AccountType:  dto.AccountType,
		AccountName:  dto.AccountName,
		CurrencyCode: dto.CurrencyCode,
		Status:       1,
		Balance:      amount,
	}
	rdto, err := domain.Create(account)
	return rdto, err
}

func (a *accountService) Transfer(dto services.AccountTransferDTO) (services.TransferStatus, error) {
	domain := accountDomain{}
	if err := base.ValidateStruct(&dto); err != nil {
		return services.TransferStatusFailure, err
	}
	amount, err := decimal.NewFromString(dto.AmountStr)
	if err != nil {
		return services.TransferStatusFailure, err
	}
	dto.Amount = amount
	if dto.ChangeFlag == services.FlagTransferOut {
		if dto.ChangeType > 0 {
			return services.TransferStatusFailure, errors.New("changeFlag不符合支出")
		}
	} else {
		if dto.ChangeType < 0 {
			return services.TransferStatusFailure, errors.New("changeFlag不符合收入")
		}
	}
	status, err := domain.Transfer(dto)
	if status == services.TransferStatusSuccess && dto.TradeBody.AccountNo != dto.TradeTarget.AccountNo && dto.ChangeType != services.AccountStoreValue {
		backwardDto := dto
		backwardDto.TradeBody = dto.TradeTarget
		backwardDto.TradeTarget = dto.TradeBody
		backwardDto.ChangeType = -dto.ChangeType
		backwardDto.ChangeFlag = -dto.ChangeFlag
		status, err := domain.Transfer(backwardDto)
		return status, err
	}
	return status, err
}

func (a *accountService) StoreValue(dto services.AccountTransferDTO) (services.TransferStatus, error) {
	dto.TradeTarget = dto.TradeBody
	dto.ChangeFlag = services.FlagTransferIn
	dto.ChangeType = services.AccountStoreValue
	return a.Transfer(dto)
}

func (a *accountService) GetEnvelopeAccountById(userId string) *services.AccountDTO {
	domain := accountDomain{}
	account := domain.GetEnvelopeAccountByUserId(userId)
	return account
}
