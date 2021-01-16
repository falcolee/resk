package base

import (
	"context"
	"github.com/jinzhu/gorm"
)

const TX = "tx"

func Tx(fn func(db *gorm.DB) error) error {
	return TxContext(context.Background(), fn)
}

//事务执行帮助函数，简化代码，需要传入上下文
func TxContext(ctx context.Context, fn func(db *gorm.DB) error) error {
	return ORM().Transaction(fn)
}
