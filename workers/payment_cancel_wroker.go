package workers

import (
	"time"
)

func StartPaymentCancelWorker() {

	for {
		println("payment time ", time.Now().String())
		time.Sleep(3 * time.Second)
	}
}
