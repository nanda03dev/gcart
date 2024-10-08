package repositories

import (
	"github.com/nanda03dev/gcart/src/config"
)

type Repositories struct {
	City          *CityRepository
	User          *UserRepository
	Product       *ProductRepository
	Order         *OrderRepository
	Item          *ItemRepository
	Payment       *PaymentRepository
	RefundPayment *RefundPaymentRepository
	Event         *EventRepository
}

var AppRepositories Repositories

func InitializeRepositories() {
	AppRepositories = Repositories{
		City:          NewCityRepository(config.DB),
		User:          NewUserRepository(config.DB),
		Product:       NewProductRepository(config.DB),
		Order:         NewOrderRepository(config.DB),
		Item:          NewItemRepository(config.DB),
		Payment:       NewPaymentRepository(config.DB),
		RefundPayment: NewRefundPaymentRepository(config.DB),
		Event:         NewEventRepository(config.DB),
	}
}
