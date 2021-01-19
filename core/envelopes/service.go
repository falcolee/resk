package envelopes

import (
	"moyutec.top/resk/services"
	"sync"
)

var once sync.Once

func init() {
	once.Do(func() {
		services.IRedEnvelopeService = new(redEnvelopeService)
	})
}

type redEnvelopeService struct {
}

func (r *redEnvelopeService) SendOut(services.RedEnvelopeSendingDTO) (activity *services.RedEnvelopeActivity, err error) {
	panic("implement me")
}

func (r *redEnvelopeService) Receive(dto services.RedEnvelopeReceiveDTO) (item *services.RedEnvelopeItemDTO, err error) {
	panic("implement me")
}

func (r *redEnvelopeService) Refund(envelopeNo string) (order *services.RedEnvelopeGoodsDTO) {
	panic("implement me")
}

func (r *redEnvelopeService) Get(envelopeNo string) (order *services.RedEnvelopeGoodsDTO) {
	domain := goodsDomain{}
	po := domain.GetOne(envelopeNo)
	if po == nil {
		return order
	}
	return po.ToDTO()
}

func (r *redEnvelopeService) ListSent(userId string, page, size int) (orders []*services.RedEnvelopeGoodsDTO) {

}

func (r *redEnvelopeService) ListReceived(userId string, page, size int) (items []*services.RedEnvelopeItemDTO) {
	panic("implement me")
}

func (r *redEnvelopeService) ListReceivable(page, size int) (orders []*services.RedEnvelopeGoodsDTO) {
	panic("implement me")
}

func (r redEnvelopeService) ListItems(envelopeNo string) (items []*services.RedEnvelopeItemDTO) {
	panic("implement me")
}
