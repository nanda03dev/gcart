package models

import (
	"github.com/nanda03dev/gnosql_client"
	"github.com/nanda03dev/go2ms/common"
	"github.com/nanda03dev/go2ms/global_constant"
)

type RefundPayment struct {
	DocId      string            `json:"docId" bson:"docId"`
	OrderId    string            `json:"orderId" bson:"orderId"`
	Name       string            `json:"name" bson:"name"`
	Amount     int               `json:"amount" bson:"amount"`
	StatusCode common.StatusCode `json:"statusCode" bson:"statusCode"`
}

var RefundPaymentGnosql = gnosql_client.CollectionInput{
	CollectionName: "refundPayments",
	IndexKeys:      []string{"orderId"},
}

func (refundPayment RefundPayment) ToDocument() gnosql_client.Document {
	return gnosql_client.Document{
		"docId":      refundPayment.DocId,
		"orderId":    refundPayment.OrderId,
		"name":       refundPayment.Name,
		"amount":     refundPayment.Amount,
		"statusCode": refundPayment.StatusCode,
	}
}

func (refundPayment RefundPayment) ToModel(refundPaymentDocument gnosql_client.Document) RefundPayment {
	return RefundPayment{
		DocId:      GetStringValue(refundPaymentDocument, "docId"),
		OrderId:    GetStringValue(refundPaymentDocument, "orderId"),
		Name:       GetStringValue(refundPaymentDocument, "name"),
		Amount:     GetIntegerValue(refundPaymentDocument, "amount"),
		StatusCode: GetValue[common.StatusCode](refundPaymentDocument, "statusCode"),
	}
}

func (refundPayment RefundPayment) ToEvent(operationType common.OperationType) common.EventType {
	return common.EventType{
		EntityId:      refundPayment.DocId,
		EntityType:    global_constant.ENTITY_REFUND_PAYMENT,
		OperationType: operationType,
		CheckProcess:  GetCheckProcess(global_constant.ENTITY_REFUND_PAYMENT, operationType),
	}
}

func (refundPayment RefundPayment) ToUpdatedDocument(newRefundPayment RefundPayment) RefundPayment {
	refundPaymentDocument := refundPayment.ToDocument()
	newRefundPaymentDocument := newRefundPayment.ToDocument()

	// statusCode should not be updated
	newRefundPaymentDocument["statusCode"] = refundPaymentDocument["statusCode"]

	for key, value := range newRefundPaymentDocument {
		if value != nil && value != "" {
			refundPaymentDocument[key] = value
		}
	}

	return refundPayment.ToModel(refundPaymentDocument)
}
