package common

var ChannelCRUD chan EventType
var ChannelPaymentRefund chan EventType

func InitializeChannels() {
	ChannelCRUD = make(chan EventType)
	ChannelPaymentRefund = make(chan EventType)
}

func AddToChanCRUD(event EventType) {
	ChannelCRUD <- event
}

func AddToChanPaymentRefund(event EventType) {
	ChannelPaymentRefund <- event
}
