package controller

import (
	"api-pedidos/core/domain"
	"api-pedidos/core/usecase/input"
	"context"
	"fmt"

	"github.com/stretchr/testify/mock"
)

type errorReader struct{}

func (e *errorReader) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("mock read error")
}

type (
	searchUserByIdMock struct {
		mock.Mock
	}
	updateStatusMock struct {
		mock.Mock
	}
)

func (uc *searchUserByIdMock) Execute(ctx *context.Context, i *input.FindByIdInput) (*[]domain.Pedido, error) {
	args := uc.Called()
	return args.Get(0).(*[]domain.Pedido), args.Error(1)
}

func (uc *updateStatusMock) Execute(ctx *context.Context, i *input.UpdateStatusInput) (*domain.Pedido, error) {
	args := uc.Called()
	return args.Get(0).(*domain.Pedido), args.Error(1)
}