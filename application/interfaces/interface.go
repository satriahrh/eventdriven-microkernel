package interfaces

import (
	"context"
	"edmk/application/core/dto"
)

type CoreService interface {
	BillerExecute(ctx context.Context, kernelID string, command string, request dto.BillerRequest) (dto.BillerResponse, error)
}
