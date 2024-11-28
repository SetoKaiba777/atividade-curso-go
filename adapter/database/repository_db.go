package database

import (
	"api-pedidos/core/domain"
	"context"
)

type DbPedidos interface {
	Save(ctx *context.Context, pedido *domain.Pedido) (domain.Pedido, error)
	FindByUserId(ctx *context.Context, uId string) ([]domain.Pedido, error)
	FindById(ctx *context.Context, id string) (domain.Pedido, error)
}