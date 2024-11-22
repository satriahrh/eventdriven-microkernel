package entity

type Kernel struct {
	ID               string
	Name             string
	Type             KernelType
	StatusDefault    KernelStatus
	StatusMachineMap map[string]KernelStatus
}

type KernelStatus uint

const (
	INACTIVE KernelStatus = iota
	ACTIVE
)

type KernelType string

const (
	EwalletTopupBillerType KernelType = "ewallet_topup_biller"
	PartnerIntegrationType KernelType = "partner_integration"
)
