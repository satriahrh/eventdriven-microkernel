package core

import (
	"context"
	"edmk/application/core/contract"
	"edmk/application/core/dto"
	"edmk/application/interfaces"
	"edmk/domain/entity"
	"edmk/domain/repository"
	"fmt"
	"os"
	"path"
	"plugin"
)

func NewCoreService() interfaces.CoreService {
	core := &core{
		kernelDirectory: make(map[string]KernelDirItem),
	}
	kernel := entity.Kernel{
		ID:            "kernela",
		Name:          "Kernel A",
		Type:          entity.EwalletTopupBillerType,
		StatusDefault: entity.ACTIVE,
	}
	if err := core.loadKernel(kernel); err != nil {
		panic(err)
	}

	return core
}

type core struct {
	kernelDirectory map[string]KernelDirItem
}

type KernelDirItem struct {
	entity.Kernel
	Implementation interface{}
}

func (c *core) loadKernel(kernel entity.Kernel) error {
	currentDir, _ := os.Getwd()
	kernelDir := path.Join(currentDir, "kernels")
	p, err := plugin.Open(fmt.Sprintf("%s/%s.so", kernelDir, kernel.ID))
	if err != nil {
		return fmt.Errorf("cannot load kernel %s: %v", kernel.ID, err)
	}

	constructSymbol, err := p.Lookup("Construct")
	if err != nil {
		return fmt.Errorf("kernel do not have Construct function")
	}

	implementation, err := c.safeConstructorCall(constructSymbol)
	if err != nil {
		return fmt.Errorf("kernel Construct function return error: %v", err)
	}

	switch kernel.Type {
	case entity.EwalletTopupBillerType:
		if _, ok := implementation.(contract.EwalletTopupBiller); !ok {
			return fmt.Errorf("kernel Construct function return not EwalletTopup")
		}
	default:
		return fmt.Errorf("kernel type %s not found", kernel.Type)
	}

	c.kernelDirectory[kernel.ID] = KernelDirItem{
		Kernel:         kernel,
		Implementation: implementation,
	}
	return nil
}

func (c *core) safeConstructorCall(constructSymbol plugin.Symbol) (impl interface{}, err error) {
	constructor, ok := constructSymbol.(func(contract.InfrastructureGetter) interface{})
	if !ok {
		return nil, fmt.Errorf("kernel Construct function is not Constructor")
	}

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("kernel Construct function panic: %v", r)
		}
	}()

	impl = constructor(func(krt repository.KernelRepositoryType) interface{} {
		return nil
	})
	return
}

func (c *core) BillerExecute(ctx context.Context, kernelID string, command string, request dto.BillerRequest) (dto.BillerResponse, error) {
	kernelDirItem, ok := c.kernelDirectory[kernelID]
	if !ok {
		return dto.BillerResponse{}, fmt.Errorf("kernel %s not found", kernelID)
	}

	switch kernelDirItem.Type {
	case entity.EwalletTopupBillerType:
		kernel := kernelDirItem.Implementation.(contract.EwalletTopupBiller)
		response, err := c.dispatchEwalletTopup(ctx, kernel, command, request)
		if err != nil {
			return dto.BillerResponse{}, err
		}
		return response, nil
		// kernel := kernelItem.Implementation.(contract.EwalletTopupBiller)
		// kernel.Advice(ctx, contract.EwalletTopupBillerPurchaseRequest{})
	case entity.PartnerIntegrationType:

	default:
		return dto.BillerResponse{}, fmt.Errorf("kernel type %s not found", kernelID)
	}
}

func (c *core) dispatchEwalletTopup(ctx context.Context, implementation contract.EwalletTopupBiller, command string, request dto.BillerRequest) (dto.BillerResponse, error) {
	switch command {
	case "inquiry":
		response, err := implementation.Inquiry(ctx, contract.EwalletTopupBillerInquiryRequest{})
		if err != nil {
			return dto.BillerResponse{}, err
		}
		return dto.BillerResponse{
			ProductID:   request.ProductID,
			CustomerID:  response.CustomerID,
			ReferenceNo: request.ReferenceNo,
			Code:        response.Code,
			Message:     response.Message,
			Amount:      0,
			Price:       0,
			AdminFee:    response.AdminFee,
			Status:      "success",
			Detail: dto.BillerResponseEwalletTopupDetail{
				CustomerID:   response.CustomerID,
				CustomerName: response.CustomerName,
				MaxAmount:    response.MaxAmount,
				MinAmount:    response.MinAmount,
			},
		}, nil

	case "purchase":
		panic("unimplemented")
	case "advice":
		panic("unimplemented")
	default:
		return dto.BillerResponse{}, fmt.Errorf("command %s not found", command)
	}
}
