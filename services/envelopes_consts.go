package services

const (
	DefaultBlessing   = "恭喜发财，鸿运牛来"
	DefaultTimeFormat = "2006-01-02.15:04:05"
)

type OrderType int

const (
	OrderTypeSending OrderType = 1
	OrderTypeRefund  OrderType = 2
)

type PayStatus int

//支付状态：未支付，支付中，已支付，支付失败
//退款：未退款，退款中，已退款，退款失败
const (
	PayNothing PayStatus = 1
	Paying     PayStatus = 2
	Payed      PayStatus = 3
	PayFailure PayStatus = 4
	//
	RefundNothing PayStatus = 61
	Refunding     PayStatus = 62
	Refunded      PayStatus = 63
	RefundFailure PayStatus = 64
)

//红包订单状态：创建、发布、过期、失效、过期退款成功，过期退款失败
type OrderStatus int

const (
	OrderCreate                  OrderStatus = 1
	OrderSending                 OrderStatus = 2
	OrderExpired                 OrderStatus = 3
	OrderDisabled                OrderStatus = 4
	OrderExpiredRefundSuccessful OrderStatus = 5
	OrderExpiredRefundFailed     OrderStatus = 6
)

//红包类型：普通红包，碰运气红包
type EnvelopeType int

const (
	GeneralEnvelopeType = 1
	LuckyEnvelopeType   = 2
)

var _ = map[EnvelopeType]string{
	GeneralEnvelopeType: "普通红包",
	LuckyEnvelopeType:   "碰运气红包",
}
