package workers

func InitiateWorker() {
	go StartCRUDWorker()
	go StartEntityTimeoutWorker()
}
