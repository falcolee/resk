package base

import (
	"github.com/jinzhu/gorm"
)

const TX = "tx"

func Tx(fn func(db *gorm.DB) error) error {
	return ORM().Transaction(fn)
}
