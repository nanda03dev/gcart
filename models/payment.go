package models

import (
	"github.com/nanda03dev/gnosql_client"
	"github.com/nanda03dev/go2ms/common"
	"github.com/nanda03dev/go2ms/global_constant"
)

type Payment struct {
	DocId      string            `json:"docId" bson:"docId"`
	OrderId    string            `json:"orderId" bson:"orderId"`
	Name       string            `json:"name" bson:"name"`
	Amount     int               `json:"amount" bson:"amount"`
	StatusCode common.StatusCode `json:"statusCode" bson:"statusCode"`
}

var PaymentGnosql = gnosql_client.CollectionInput{
	CollectionName: "payments",
	IndexKeys:      []string{"orderId"},
}

func (payment Payment) ToDocument() gnosql_client.Document {
	return gnosql_client.Document{
		"docId":      payment.DocId,
		"orderId":    payment.OrderId,
		"name":       payment.Name,
		"amount":     payment.Amount,
		"statusCode": payment.StatusCode,
	}
}

func (payment Payment) ToModel(paymentDocument gnosql_client.Document) Payment {
	return Payment{
		DocId:      GetStringValue(paymentDocument, "docId"),
		OrderId:    GetStringValue(paymentDocument, "orderId"),
		Name:       GetStringValue(paymentDocument, "name"),
		Amount:     GetIntegerValue(paymentDocument, "amount"),
		StatusCode: GetValue[common.StatusCode](paymentDocument, "statusCode"),
	}
}

func (payment Payment) ToEvent(operationType common.OperationType) common.EventType {
	return common.EventType{
		EntityId:      payment.DocId,
		EntityType:    global_constant.ENTITY_PAYMENT,
		OperationType: operationType,
		CheckProcess:  GetCheckProcess(global_constant.ENTITY_PAYMENT, operationType),
	}
}

func (payment Payment) ToUpdatedDocument(newPayment Payment) Payment {
	paymentDocument := payment.ToDocument()
	newPaymentDocument := newPayment.ToDocument()

	// statusCode should not be updated
	newPaymentDocument["statusCode"] = paymentDocument["statusCode"]

	for key, value := range newPaymentDocument {
		if value != nil && value != "" {
			paymentDocument[key] = value
		}
	}

	return payment.ToModel(paymentDocument)
}
