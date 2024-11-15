package usecase

import (
	"api-pedidos/core/domain"
	"context"

	"github.com/stretchr/testify/mock"
)

type databaseMock struct {
	mock.Mock
}

func NewMockDatabase() *databaseMock {
	return &databaseMock{}
}

func (d *databaseMock) Add(ctx *context.Context, p *domain.Pedido) (domain.Pedido, error) {
	args := d.Called()
	return args.Get(0).(domain.Pedido), args.Error(1)
}
