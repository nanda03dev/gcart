package workers

func InitiateWorker() {
	go StartOrderWorker()
	go StartPaymentCancelWorker()
}
