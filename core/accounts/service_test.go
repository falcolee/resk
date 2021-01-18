package accounts

import (
	"github.com/segmentio/ksuid"
	"moyutec.top/resk/services"
	"reflect"
	"testing"
)

func Test_accountService_CreateAccount(t *testing.T) {
	dto := services.AccountCreatedDTO{
		UserId:       ksuid.New().Next().String(),
		Username:     "测试用户a",
		Amount:       "100",
		AccountName:  "测试账户a",
		AccountType:  1,
		CurrencyCode: "CNY",
	}
	type args struct {
		dto services.AccountCreatedDTO
	}
	tests := []struct {
		name    string
		args    args
		want    *services.AccountDTO
		wantErr bool
	}{
		{name: "test1", args: args{dto: dto}, want: nil, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &accountService{}
			got, err := a.CreateAccount(tt.args.dto)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				if got.Username != dto.Username {
					t.Errorf("CreateAccount() got = %v", got)
				}
			}
		})
	}
}

func Test_accountService_GetEnvelopeAccountById(t *testing.T) {
	a1 := services.AccountCreatedDTO{
		UserId:       ksuid.New().Next().String(),
		Username:     "测试用户1",
		Amount:       "100",
		AccountName:  "测试账户1",
		AccountType:  1,
		CurrencyCode: "CNY",
	}
	service := &accountService{}
	adto1, err := service.CreateAccount(a1)
	if err != nil {
		t.Errorf("CreateAccount() error = %v", err)
	}
	type args struct {
		userId string
	}
	tests := []struct {
		name string
		args args
		want *services.AccountDTO
	}{
		{name: "test1", args: args{userId: adto1.UserId}, want: adto1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := service.GetEnvelopeAccountById(tt.args.userId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEnvelopeAccountById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_accountService_Transfer(t *testing.T) {
	//准备2个账户
	a1 := services.AccountCreatedDTO{
		UserId:       ksuid.New().Next().String(),
		Username:     "测试用户1",
		Amount:       "100",
		AccountName:  "测试账户1",
		AccountType:  2,
		CurrencyCode: "CNY",
	}
	a2 := services.AccountCreatedDTO{
		UserId:       ksuid.New().Next().String(),
		Username:     "测试用户2",
		Amount:       "100",
		AccountName:  "测试账户2",
		AccountType:  2,
		CurrencyCode: "CNY",
	}
	service := &accountService{}
	adto1, err := service.CreateAccount(a1)
	if err != nil {
		t.Errorf("CreateAccount1() error = %v", err)
	}
	if adto1 == nil {
		t.Errorf("CreateAccount1() error = %v", err)
	}
	adto2, err := service.CreateAccount(a2)
	if err != nil {
		t.Errorf("CreateAccount2() error = %v", err)
	}
	if adto2 == nil {
		t.Errorf("CreateAccount2() error = %v", err)
	}
	body := services.TradeParticipator{
		AccountNo: adto1.AccountNo,
		UserId:    adto1.UserId,
		Username:  adto1.Username,
	}
	target := body
	type args struct {
		dto services.AccountTransferDTO
	}
	tests := []struct {
		name    string
		args    args
		want    services.TransferStatus
		wantErr bool
	}{
		{name: "账户1储值", args: args{dto: services.AccountTransferDTO{
			TradeBody:   body,
			TradeTarget: target,
			TradeNo:     ksuid.New().Next().String(),
			AmountStr:   "10",
			ChangeType:  services.AccountStoreValue,
			ChangeFlag:  services.FlagTransferIn,
			Decs:        "储值1",
		}}, want: services.TransferStatusSuccess, wantErr: false},
		{name: "账户2储值", args: args{dto: services.AccountTransferDTO{
			TradeBody: body,
			TradeTarget: services.TradeParticipator{
				AccountNo: adto2.AccountNo,
				UserId:    adto2.UserId,
				Username:  adto2.Username},
			TradeNo:    ksuid.New().Next().String(),
			AmountStr:  "20",
			ChangeType: services.ChangeType(-1),
			ChangeFlag: services.FlagTransferOut,
			Decs:       "转出",
		}}, want: services.TransferStatusSuccess, wantErr: false},
		{name: "账户2储值", args: args{dto: services.AccountTransferDTO{
			TradeBody: body,
			TradeTarget: services.TradeParticipator{
				AccountNo: adto2.AccountNo,
				UserId:    adto2.UserId,
				Username:  adto2.Username},
			TradeNo:    ksuid.New().Next().String(),
			AmountStr:  "120",
			ChangeType: services.ChangeType(-1),
			ChangeFlag: services.FlagTransferOut,
			Decs:       "转出",
		}}, want: services.TransferStatusFundsNotAllowed, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := service.Transfer(tt.args.dto)
			if (err != nil) != tt.wantErr {
				t.Errorf("Transfer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Transfer() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_accountService_GetAccount(t *testing.T) {
	a1 := services.AccountCreatedDTO{
		UserId:       ksuid.New().Next().String(),
		Username:     "测试用户1",
		Amount:       "100",
		AccountName:  "测试账户1",
		AccountType:  2,
		CurrencyCode: "CNY",
	}
	service := &accountService{}
	adto1, err := service.CreateAccount(a1)
	if err != nil {
		t.Errorf("CreateAccount() error = %v", err)
	}
	type args struct {
		accountNo string
	}
	tests := []struct {
		name string
		args args
		want *services.AccountDTO
	}{
		{name: "test1", args: args{accountNo: adto1.AccountNo}, want: adto1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := service.GetAccount(tt.args.accountNo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}
