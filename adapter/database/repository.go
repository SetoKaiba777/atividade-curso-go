package database

import (
	"api-pedidos/core/domain"
	"api-pedidos/core/repository"
	"context"
)

type PedidosRepository struct {
	db DbPedidos
}

var _ repository.PedidoRepository = (*PedidosRepository)(nil)

func (r *PedidosRepository) Add(ctx *context.Context, pedido *domain.Pedido) (domain.Pedido, error){
	p, err := r.db.Save(ctx, pedido)
	if err != nil{
		return p,err
	}
	return p, nil
}
func (r *PedidosRepository) GetByUserId(ctx *context.Context, uId string) ([]domain.Pedido, error){
	ps, err := r.db.FindByUserId(ctx,uId)
	if err != nil{
		return ps, err
	}
	return ps, nil
}
func (r *PedidosRepository) GetById(ctx *context.Context, id string) (domain.Pedido, error){
	p, err := r.db.FindById(ctx, id)
	if err != nil{
		return p, err
	}
	return p, nil
}