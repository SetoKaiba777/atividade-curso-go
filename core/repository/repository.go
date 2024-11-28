package repository

import (
	"api-pedidos/core/domain"
	"context"
)

type PedidoRepository interface {
	Add(ctx *context.Context, pedido *domain.Pedido) (domain.Pedido, error)
	GetByUserId(ctx *context.Context, uId string) ([]domain.Pedido, error)
	GetById(ctx *context.Context, id string) (domain.Pedido, error)
}
