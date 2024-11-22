package repository

import (
	"edmk/domain/entity"
	"io"
	"net/http"
)

type KernelRepositoryType uint

const (
	DURABLE_STORAGE KernelRepositoryType = iota
	TEMPORARY_STORAGE
	KERNEL_QUEUE
	HTTP_CLIENT
)

type DurableStorageRepository interface {
	Store(data string)
}

type TemporaryStorageRepository interface {
	Store(data string)
	Get() string
}

type KernelQueueRepository interface {
}

type HttpClientRepository interface {
	CloseIdleConnections()
	Do(*http.Request) (*http.Response, error)
	Get(url string) (*http.Response, error)
	Post(url string, contentType string, body io.Reader) (*http.Response, error)
	PostForm(url string, data map[string][]string) (*http.Response, error)
}

type KernelRegistry interface {
	GetAllKernels() ([]entity.Kernel, error)
	GetByID(id string) (entity.Kernel, error)
	UpdateStatus(id string, machine string, status entity.KernelStatus) error
}
