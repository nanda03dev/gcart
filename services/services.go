package services

import (
	"github.com/nanda03dev/gcart/config"
	"github.com/nanda03dev/gcart/repositories"
)

type Services struct {
	City          CityService
	User          UserService
	Order         OrderService
	Item          ItemService
	Product       ProductService
	Payment       PaymentService
	RefundPayment RefundPaymentService
	Event         EventService
}

var AppServices Services

func InitializeServices() {
	AppServices = Services{
		City:          NewCityService(repositories.NewCityRepository(config.DB)),
		User:          NewUserService(repositories.NewUserRepository(config.DB)),
		Order:         NewOrderService(repositories.NewOrderRepository(config.DB)),
		Item:          NewItemService(repositories.NewItemRepository(config.DB)),
		Product:       NewProductService(repositories.NewProductRepository(config.DB)),
		Payment:       NewPaymentService(repositories.NewPaymentRepository(config.DB)),
		RefundPayment: NewRefundPaymentService(repositories.NewRefundPaymentRepository(config.DB)),
		Event:         NewEventService(repositories.NewEventRepository(config.DB)),
	}
}
