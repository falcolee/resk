package accounts

import (
	"github.com/segmentio/ksuid"
	"github.com/shopspring/decimal"
	"moyutec.top/resk/infra/base"
	_ "moyutec.top/resk/testx"
	"testing"
)

func TestAccountDao_GetByUserId(t *testing.T) {
	type args struct {
		userId      string
		accountType int
	}
	runner := base.ORM()
	a := &Account{
		UserId:   "1nEQTTh38lLvzDRUWYpCWVow3cv",
		Username: "测试用户",
	}
	tests := []struct {
		name string
		args args
		want *Account
	}{
		{name: "空", args: struct {
			userId      string
			accountType int
		}{userId: "xxxxxxx", accountType: 1}, want: nil},
		{name: "非空", args: struct {
			userId      string
			accountType int
		}{userId: "1nEQTTh38lLvzDRUWYpCWVow3cv", accountType: 2}, want: a},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dao := &AccountDao{
				runner: runner,
			}
			got := dao.GetByUserId(tt.args.userId, tt.args.accountType)
			if (tt.want != nil) && (tt.want.UserId != got.UserId && tt.want.Username != got.Username) {
				t.Errorf("GetByUserId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountDao_GetOne(t *testing.T) {
	runner := base.ORM()
	type args struct {
		accountNo string
	}
	a := &Account{
		Balance:      decimal.NewFromFloat(100),
		Status:       1,
		AccountNo:    "1nEQTT6VFPiP2t1W4nCVp90q4yf",
		AccountName:  "测试资金账户",
		Username:     "测试用户",
		CurrencyCode: "CNY",
		AccountType:  2,
	}
	tests := []struct {
		name string
		args args
		want *Account
	}{
		{name: "空", args: struct {
			accountNo string
		}{accountNo: "xxxxxxx"}, want: nil},
		{name: "非空", args: struct {
			accountNo string
		}{accountNo: "1nEQTT6VFPiP2t1W4nCVp90q4yf"}, want: a},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dao := &AccountDao{
				runner: runner,
			}
			got := dao.GetOne(tt.args.accountNo)
			if (tt.want != nil) && (tt.want.AccountNo != got.AccountNo && tt.want.Username != got.Username) {
				t.Errorf("GetOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountDao_Insert(t *testing.T) {
	runner := base.ORM()
	a := &Account{
		Balance:      decimal.NewFromFloat(100),
		Status:       1,
		AccountNo:    ksuid.New().Next().String(),
		AccountName:  "测试资金账户",
		UserId:       ksuid.New().Next().String(),
		Username:     "测试用户",
		CurrencyCode: "CNY",
		AccountType:  2,
	}
	type args struct {
		account *Account
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "测试用户", args: struct{ account *Account }{account: a}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dao := &AccountDao{
				runner: runner,
			}
			gotId, err := dao.Insert(tt.args.account)
			if (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId <= 0 {
				t.Errorf("Insert() gotId = %v, want >0", gotId)
			}
		})
	}
}
