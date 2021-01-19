package envelopes

import (
	"github.com/jinzhu/gorm"
	"github.com/segmentio/ksuid"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"moyutec.top/resk/infra/base"
	"moyutec.top/resk/services"
	"time"
)

type goodsDomain struct {
	RedEnvelopeGoods
	item itemDomain
}

func (d *goodsDomain) createEnvelopeNo() {
	d.EnvelopeNo = ksuid.New().Next().String()
}

func (d *goodsDomain) Create(goods services.RedEnvelopeGoodsDTO) {
	d.RedEnvelopeGoods.FromDTO(&goods)
	d.RemainQuantity = goods.Quantity
	if d.EnvelopeType == services.GeneralEnvelopeType {
		d.Amount = goods.AmountOne.Mul(decimal.NewFromFloat(float64(goods.Quantity)))
	}
	if d.EnvelopeType == services.LuckyEnvelopeType {
		d.AmountOne = decimal.NewFromFloat(0)
	}
	d.RemainAmount = d.Amount
	d.ExpiredAt = time.Now().Add(24 * time.Hour)
	d.Status = services.OrderCreate
	d.OrderType = services.OrderTypeSending
	d.PayStatus = services.Paying
	d.createEnvelopeNo()
}

func (d *goodsDomain) Save() (id int64, err error) {
	err = base.Tx(func(runner *gorm.DB) error {
		dao := RedEnvelopeGoodsDao{runner: runner}
		id, err = dao.Insert(&d.RedEnvelopeGoods)
		return err
	})
	if err != nil {
		logrus.Error(err)
	}
	return id, err
}

func (d *goodsDomain) CreateAndSave(goods services.RedEnvelopeGoodsDTO) (id int64, err error) {
	//创建红包商品
	d.Create(goods)
	//保存红包商品
	return d.Save()
}

func (d *goodsDomain) GetOne(envelopeNo string) (goods *RedEnvelopeGoods) {
	dao := RedEnvelopeGoodsDao{runner: base.ORM()}
	goods = dao.GetOne(envelopeNo)
	return goods
}

func (d *goodsDomain) UpdateOrderStatus(envelopeNo string, status services.OrderStatus) (rows int64, err error) {
	err = base.Tx(func(runner *gorm.DB) error {
		dao := RedEnvelopeGoodsDao{runner: runner}
		rows, err = dao.UpdateOrderStatus(envelopeNo, status)
		return nil
	})
	if err != nil {
		logrus.Error(err)
	}
	return
}

func (d *goodsDomain) UpdatePayStatus(envelopeNo string, status services.PayStatus) (rows int64, err error) {
	err = base.Tx(func(runner *gorm.DB) error {
		dao := RedEnvelopeGoodsDao{runner: runner}
		rows, err = dao.UpdatePayStatus(envelopeNo, status)
		return nil
	})
	if err != nil {
		logrus.Error(err)
	}
	return
}
