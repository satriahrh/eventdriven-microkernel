package entity

type Product struct {
	ID       string
	Name     string
	Operator string
	Nominal  int
	Pricing  Pricing
}

type Pricing struct {
	Type   string
	Amount int
}
