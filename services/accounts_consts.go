package services

type TransferStatus int8

//转账状态
const (
	//失败
	TransferStatusFailure TransferStatus = -1
	//余额不足
	TransferStatusFundsNotAllowed TransferStatus = 0
	TransferStatusSuccess         TransferStatus = 1
)

//转账类型
type ChangeType int8

const (
	//创建账户
	AccountCreated ChangeType = 0
	//储值
	AccountStoreValue ChangeType = 1
	//红包资金的支出、收入、过期退款
	EnvelopeOutgoing ChangeType = -2
	//收入
	EnvelopeIncoming ChangeType = 2
	//过期退款
	EnvelopeExpiredRefund ChangeType = 3
)

//资金交易的变化标识
type ChangeFlag int8

const (
	//创建账户=0
	//支出=-1
	//收入=2
	FlagAccountCreated ChangeFlag = 0
	FlagTransferOut    ChangeFlag = -1
	FlagTransferIn     ChangeFlag = 1
)
