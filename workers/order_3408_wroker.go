package workers

import (
	// "context"
	"time"

	// "github.com/nanda03dev/go2ms/repositories"
)

func Start3408Worker() {

	for {
		// repositories.AppRepositories.Event.GetAll(context.TODO(), nil, nil, nil)
		time.Sleep(3 * time.Second)
	}
}
