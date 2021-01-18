package accounts

import (
	"github.com/segmentio/ksuid"
	"github.com/shopspring/decimal"
	"moyutec.top/resk/infra/base"
	"moyutec.top/resk/services"
	_ "moyutec.top/resk/testx"
	"testing"
)

func TestAccountLogDao_GetByTradeNo(t *testing.T) {
	runner := base.ORM()
	type args struct {
		tradeNo string
	}
	a := &AccountLog{
		TradeNo:  "1nESLccMbWRDuhB2zVSWhIL5tJp",
		LogNo:    "1nESLaFIz4tsOkDnhMO79SOetP5",
		Username: "测试用户",
	}
	tests := []struct {
		name string
		args args
		want *AccountLog
	}{
		{name: "空", args: args{tradeNo: "xxxx"}, want: nil},
		{name: "非空", args: args{tradeNo: "1nESLccMbWRDuhB2zVSWhIL5tJp"}, want: a},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dao := &AccountLogDao{
				runner: runner,
			}
			got := dao.GetByTradeNo(tt.args.tradeNo)
			if tt.want != nil && (tt.want.TradeNo != got.TradeNo || tt.want.Username != got.Username) {
				t.Errorf("GetByTradeNo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountLogDao_GetOne(t *testing.T) {
	runner := base.ORM()
	type args struct {
		logNo string
	}
	a := &AccountLog{
		LogNo:    "1nESLaFIz4tsOkDnhMO79SOetP5",
		Username: "测试用户",
	}
	tests := []struct {
		name string
		args args
		want *AccountLog
	}{
		{name: "空", args: args{logNo: "xxx"}, want: nil},
		{name: "非空", args: args{logNo: "1nESLaFIz4tsOkDnhMO79SOetP5"}, want: a},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dao := &AccountLogDao{
				runner: runner,
			}
			got := dao.GetOne(tt.args.logNo)
			if tt.want != nil && (tt.want.LogNo != got.LogNo || tt.want.Username != got.Username) {
				t.Errorf("GetOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountLogDao_Insert(t *testing.T) {
	runner := base.ORM()
	a := &AccountLog{
		LogNo:      ksuid.New().Next().String(),
		TradeNo:    ksuid.New().Next().String(),
		Status:     1,
		AccountNo:  ksuid.New().Next().String(),
		UserId:     ksuid.New().Next().String(),
		Username:   "测试用户",
		Amount:     decimal.NewFromFloat(1),
		Balance:    decimal.NewFromFloat(100),
		ChangeFlag: services.FlagAccountCreated,
		ChangeType: services.AccountCreated,
	}
	type args struct {
		accountLog *AccountLog
	}
	tests := []struct {
		name    string
		args    args
		wantId  int64
		wantErr bool
	}{
		{name: "空", args: args{a}, wantId: 0, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dao := &AccountLogDao{
				runner: runner,
			}
			gotId, err := dao.Insert(tt.args.accountLog)
			if (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId <= tt.wantId {
				t.Errorf("Insert() gotId = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}
