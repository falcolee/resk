package envelopes

import (
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"moyutec.top/resk/services"
	"time"
)

type RedEnvelopeGoodsDao struct {
	runner *gorm.DB
}

func (dao *RedEnvelopeGoodsDao) Insert(po *RedEnvelopeGoods) (int64, error) {
	result := dao.runner.Create(po)
	err := result.Error
	return po.Id, err
}

func (dao *RedEnvelopeGoodsDao) GetOne(envelopeNo string) *RedEnvelopeGoods {
	po := &RedEnvelopeGoods{EnvelopeNo: envelopeNo}
	result := dao.runner.Where(po).First(po)
	if result.RecordNotFound() {
		return nil
	} else if result.Error != nil {
		logrus.Error(result.Error)
		return nil
	}
	return po
}

func (dao *RedEnvelopeGoodsDao) UpdateBalance(envelopeNo string, amount decimal.Decimal) (int64, error) {
	sql := "update red_envelope_goods " +
		" set remain_amount=remain_amount-CAST(? AS DECIMAL(30,6)), " +
		" remain_quantity=remain_quantity-1 " +
		" where envelope_no=? " +
		//最重要的，乐观锁的关键
		" and remain_quantity>0" +
		" and remain_amount >= CAST(? AS DECIMAL(30,6)) "
	result := dao.runner.Exec(sql, amount.String(), envelopeNo, amount.String())
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (dao *RedEnvelopeGoodsDao) UpdateOrderStatus(envelopeNo string, status services.OrderStatus) (int64, error) {
	po := &RedEnvelopeGoods{EnvelopeNo: envelopeNo}
	rs := dao.runner.Where(po).Update("status", status)
	if rs.Error != nil {
		return 0, rs.Error
	}
	return rs.RowsAffected, nil
}

func (dao *RedEnvelopeGoodsDao) UpdatePayStatus(envelopeNo string, status services.PayStatus) (int64, error) {
	po := &RedEnvelopeGoods{EnvelopeNo: envelopeNo}
	rs := dao.runner.Where(po).Update("pay_status", status)
	if rs.Error != nil {
		return 0, rs.Error
	}
	return rs.RowsAffected, nil
}

func (dao *RedEnvelopeGoodsDao) FindExpired(offset, size int) []RedEnvelopeGoods {
	var goods []RedEnvelopeGoods
	now := time.Now()
	rs := dao.runner.Model(&RedEnvelopeGoods{}).
		Where("remain_quantity > ?", 0).
		Where("order_type=?", 1).
		Where("expired_at>?", now).
		Where("status IN (?)", []int{1, 2, 3, 6}).Limit(size).Offset(offset).Find(&goods)
	if rs.Error != nil {
		logrus.Error(rs.Error)
	}
	return goods
}

func (dao *RedEnvelopeGoodsDao) FindByUser(userId string, offset, size int) []RedEnvelopeGoods {
	var goods []RedEnvelopeGoods
	rs := dao.runner.Model(&RedEnvelopeGoods{}).Where("user_id=?", userId).Limit(size).Offset(offset).Find(&goods)
	if rs.Error != nil {
		logrus.Error(rs.Error)
	}
	return goods
}

func (dao *RedEnvelopeGoodsDao) ListReceivable(offset, size int) []RedEnvelopeGoods {
	var goods []RedEnvelopeGoods
	now := time.Now()
	rs := dao.runner.Model(&RedEnvelopeGoods{}).Where("expired_at > ?", now).
		Where("remain_quantity > ?", 0).Order("created_at desc").
		Limit(size).Offset(offset).Find(&goods)
	if rs.Error != nil {
		logrus.Error(rs.Error)
	}
	return goods
}
