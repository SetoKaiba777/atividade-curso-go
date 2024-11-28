package usecase

import (
	"api-pedidos/core/domain"
	"context"
	"fmt"

	"github.com/stretchr/testify/mock"
)

type databaseMock struct {
	mock.Mock
}

var databaseErr = fmt.Errorf("Database error")

func NewMockDatabase() *databaseMock {
	return &databaseMock{}
}

func (d *databaseMock) Add(ctx *context.Context, p *domain.Pedido) (domain.Pedido, error) {
	args := d.Called()
	return args.Get(0).(domain.Pedido), args.Error(1)
}

func (d *databaseMock) GetByUserId(ctx *context.Context, uId string) ([]domain.Pedido, error) {
	args := d.Called()
	return args.Get(0).([]domain.Pedido), args.Error(1)
}

func (d *databaseMock) GetById(ctx *context.Context, uId string) (domain.Pedido, error) {
	args := d.Called()
	return args.Get(0).(domain.Pedido), args.Error(1)
}
