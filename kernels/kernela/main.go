package main

import (
	"context"
	"edmk/application/core/contract"
	"edmk/domain/repository"
)

func main() {
	// kernel := Construct(nil)
}

func Construct(infrastructureGetter contract.InfrastructureGetter) interface{} {
	// bikin queue
	// bikin http client
	// secrets := infrastructureGetter("secret")
	// secretkernela := secrets.Get("kernela")

	return &kernela{
		config: "/urls/sdfjdslk",
		// secret:     secretkernela,
		httpClient: infrastructureGetter(repository.HTTP_CLIENT).(repository.HttpClientRepository),
	}
}

type kernela struct {
	config     string
	httpClient repository.HttpClientRepository
}

// Inquiry implements contract.EwalletTopup.
func (k *kernela) Inquiry(context.Context, contract.EwalletTopupBillerInquiryRequest) (contract.EwalletTopupBillerInquiryResponse, error) {
	return contract.EwalletTopupBillerInquiryResponse{
		CustomerID:   "12",
		CustomerName: "John Doe",
		MinAmount:    10000,
		MaxAmount:    100000,
		Code:         "00",
		Message:      "Success",
		AdminFee:     2000,
	}, nil
}

// Purchase implements contract.EwalletTopup.
func (k *kernela) Purchase(context.Context, contract.EwalletTopupBillerPurchaseRequest) (contract.EwalletTopupBillerPurchaseResponse, error) {
	panic("unimplemented")
}

// Advice implements contract.EwalletTopup.
func (k *kernela) Advice(context.Context, contract.EwalletTopupBillerPurchaseRequest) (contract.EwalletTopupBillerPurchaseResponse, error) {
	panic("unimplemented")
}

func (k *kernela) Worker() {
	consumerCounter := 4
}
