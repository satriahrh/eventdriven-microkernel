package contract

import "edmk/domain/repository"

// InfrastructureGetter is a function type that returns an infrastructure implementation.
// Calling this function requires panic recovery.
type InfrastructureGetter func(repository.KernelRepositoryType) interface{}

// Constructor is a function type that returns a kernel implementation.
// Calling this function requires panic recovery.
type Constructor func(InfrastructureGetter) interface{}
