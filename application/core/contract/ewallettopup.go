package contract

import (
	"context"
	"edmk/domain/entity"
)

type EwalletTopupBillerInquiryRequest struct {
	CustomerID string
	Amount     int
	Product    entity.Product
}

type EwalletTopupBillerInquiryResponse struct {
	CustomerID   string
	CustomerName string
	MaxAmount    int
	MinAmount    int
	AdminFee     int
	Code         string
	Message      string
}

type EwalletTopupBillerPurchaseRequest struct {
	TransactionID string
	CustomerID    string
	Amount        int
	Product       entity.Product
}

type EwalletTopupBillerPurchaseResponse struct {
	TransactionID string
	CustomerID    string
	CustomerName  string
	Amount        int
	Price         int
	Status        string
	Code          string
	Message       string
}

type EwalletTopupBiller interface {
	Inquiry(context.Context, EwalletTopupBillerInquiryRequest) (EwalletTopupBillerInquiryResponse, error)
	Purchase(context.Context, EwalletTopupBillerPurchaseRequest) (EwalletTopupBillerPurchaseResponse, error)
	Advice(context.Context, EwalletTopupBillerPurchaseRequest) (EwalletTopupBillerPurchaseResponse, error)
}

type PartnerIntegration interface {
	Execute()
	Callback()
}
