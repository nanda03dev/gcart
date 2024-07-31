package common

var ChannelCRUD chan EventType

func InitializeChannels() {
	ChannelCRUD = make(chan EventType)
}

func AddToChanCRUD(event EventType) {
	ChannelCRUD <- event
}
