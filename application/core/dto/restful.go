package dto

type BillerRequest struct {
	ProductID   string
	CustomerID  string
	ReferenceNo string
	Amount      int
}

type BillerResponse struct {
	TransactionID string
	ProductID     string
	CustomerID    string
	ReferenceNo   string
	Amount        int
	Price         int
	AdminFee      int
	Status        string
	Code          string
	Message       string
	Detail        interface{}
}

type BillerResponseEwalletTopupDetail struct {
	CustomerID   string
	CustomerName string
	MaxAmount    int
	MinAmount    int
}
