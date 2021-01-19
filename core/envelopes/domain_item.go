package envelopes

import (
	"github.com/jinzhu/gorm"
	"github.com/segmentio/ksuid"
	"moyutec.top/resk/infra/base"
	"moyutec.top/resk/services"
)

type itemDomain struct {
	RedEnvelopeItem
}

func (d *itemDomain) createItemNo() {
	d.ItemNo = ksuid.New().Next().String()
}

func (d *itemDomain) Create(item services.RedEnvelopeItemDTO) {
	d.RedEnvelopeItem.FromDTO(&item)
	d.createItemNo()
}

func (d *itemDomain) Save() (id int64, err error) {
	err = base.Tx(func(runner *gorm.DB) error {
		dao := RedEnvelopeItemDao{runner: runner}
		id, err = dao.Insert(&d.RedEnvelopeItem)
		return err
	})
	return id, err
}

func (d *itemDomain) GetOne(itemNo string) (dto *services.RedEnvelopeItemDTO) {
	dao := RedEnvelopeItemDao{runner: base.ORM()}
	po := dao.GetOne(itemNo)
	if po == nil {
		return nil
	}
	dto = po.ToDTO()
	return dto
}

func (d *itemDomain) GetByUser(userId, envelopNo string) (dto *services.RedEnvelopeItemDTO) {
	dao := RedEnvelopeItemDao{runner: base.ORM()}
	po := dao.GetByUser(envelopNo, userId)
	if po == nil {
		return nil
	}
	dto = po.ToDTO()
	return dto
}

func (d *itemDomain) FindItems(envelopeNo string) (itemDtos []*services.RedEnvelopeItemDTO) {
	dao := RedEnvelopeItemDao{runner: base.ORM()}
	items := dao.FindItems(envelopeNo)
	itemDtos = make([]*services.RedEnvelopeItemDTO, 0)
	if len(items) == 0 {
		return itemDtos
	}
	var luckItem *services.RedEnvelopeItemDTO

	for i, po := range items {
		if po == nil {
			continue
		}
		item := po.ToDTO()

		if i == 0 {
			luckItem = item
		} else {
			if luckItem != nil && luckItem.Amount.Cmp(po.Amount) < 0 {
				luckItem = item
			}
		}
		itemDtos = append(itemDtos, item)
	}
	if luckItem != nil {
		luckItem.IsLuckiest = true
	}
	return itemDtos
}
