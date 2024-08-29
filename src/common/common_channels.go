package common

var ChannelCRUD chan EventType
var ChannelPaymentRefund chan EventType

func InitializeChannels() {
	ChannelCRUD = make(chan EventType, 10000)
	ChannelPaymentRefund = make(chan EventType, 10000)
}

func AddToChanCRUD(event EventType) {
	ChannelCRUD <- event
}

func AddToChanPaymentRefund(event EventType) {
	ChannelPaymentRefund <- event
}
