package workers

import (
	"fmt"
	"time"
)

func InitializeWorker() {
	go startWorkerWithRecovery(StartCRUDWorker)
	go startWorkerWithRecovery(StartEntityTimeoutWorker)
	go startWorkerWithRecovery(StartPaymentRefundWorker)
}

func startWorkerWithRecovery(workerFunc func()) {
	for {
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("Worker panic recovered: %v. Restarting worker...\n", r)
				}
			}()

			workerFunc()
		}()

		// Optional: Add a small delay before restarting the worker
		time.Sleep(2 * time.Second)
	}
}
