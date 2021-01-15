package services

type AccountService interface {
	CreateAccount(dto AccountCreatedDTO) (*AccountDTO, error)
	Transfer(dto AccountTransferDTO) (TransferStatus, error)
	StoreValue(dto AccountTransferDTO) (TransferStatus, error)
	GetEnvelopeAccountById(userId string) *AccountDTO
}

//账户交易的参与者
type TradeParticipator struct {
	AccountNo string
	UserId    string
	Username  string
}

//账户转账
type AccountTransferDTO struct {
	TradeNo     string
	TradeBody   TradeParticipator
	TradeTarget TradeParticipator
	AmountStr   string
	ChangeType  ChangeType
	ChangeFlag  ChangeFlag
	Desc        string
}

type AccountCreatedDTO struct {
	UserId       string
	Username     string
	AccountName  string
	AccountType  int
	CurrencyCode string
	Amount       string
}

type AccountDTO struct {
	AccountCreatedDTO
	AccountNo string
}
