package workers

func InitiateWorker() {
	go Start3201Worker()
	go Start3408Worker()
}
