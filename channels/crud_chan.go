package channels

import "github.com/nanda03dev/go2ms/common"

var ChannelCRUD = make(chan common.EventType)

func AddToChanCRUD(event common.EventType) {
	ChannelCRUD <- event
}
