package envelopes

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type RedEnvelopeItemDao struct {
	runner *gorm.DB
}

func (dao *RedEnvelopeItemDao) GetOne(itemNo string) *RedEnvelopeItem {
	from := &RedEnvelopeItem{ItemNo: itemNo}
	result := dao.runner.Where(from).First(from)
	if result.RecordNotFound() {
		return nil
	} else if result.Error != nil {
		logrus.Error(result.Error)
		return nil
	}
	return from
}

func (dao *RedEnvelopeItemDao) Insert(po *RedEnvelopeItem) (int64, error) {
	result := dao.runner.Create(po)
	err := result.Error
	return po.Id, err
}

func (dao *RedEnvelopeItemDao) FindItems(envelopeNo string) []*RedEnvelopeItem {
	from := &RedEnvelopeItem{EnvelopeNo: envelopeNo}
	items := make([]*RedEnvelopeItem, 0)
	rs := dao.runner.Where(from).Find(&items)
	if rs.Error != nil {
		logrus.Error(rs.Error)
	}
	return items
}

func (dao *RedEnvelopeItemDao) GetByUser(envelopeNo, userId string) *RedEnvelopeItem {
	from := &RedEnvelopeItem{EnvelopeNo: envelopeNo, RecvUserId: userId}
	result := dao.runner.Where(from).First(from)
	if result.RecordNotFound() {
		return nil
	} else if result.Error != nil {
		logrus.Error(result.Error)
		return nil
	}
	return from
}

func (dao *RedEnvelopeItemDao) ListReceivedItems(userId string, page, size int) []*RedEnvelopeItem {
	from := &RedEnvelopeItem{RecvUserId: userId}
	items := make([]*RedEnvelopeItem, 0)
	offset := (page - 1) * size
	rs := dao.runner.Where(from).Order("created_at desc").Limit(size).Offset(offset).Find(&items)
	if rs.Error != nil {
		logrus.Error(rs.Error)
	}
	return items
}
