package repositories

import (
	"github.com/nanda03dev/go2ms/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repositories struct {
	City    *CityRepository
	User    *UserRepository
	Order   *OrderRepository
	Item    *ItemRepository
	Product *ProductRepository
	Payment *PaymentRepository
	Event   *EventRepository
}

var AppRepositories Repositories

func InitializeRepositories(db *mongo.Database) {
	AppRepositories = Repositories{
		City:    NewCityRepository(config.DB),
		User:    NewUserRepository(config.DB),
		Order:   NewOrderRepository(config.DB),
		Item:    NewItemRepository(config.DB),
		Product: NewProductRepository(config.DB),
		Payment: NewPaymentRepository(config.DB),
		Event:   NewEventRepository(config.DB),
	}
}
