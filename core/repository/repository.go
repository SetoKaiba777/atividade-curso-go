package repository

import (
	"api-pedidos/core/domain"
	"context"
)

type PedidoRepository interface {
	Add(ctx *context.Context,pedido *domain.Pedido) (domain.Pedido,error)
}