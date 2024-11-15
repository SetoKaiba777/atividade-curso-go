package usecase

import (
	"api-pedidos/core/domain"
	"api-pedidos/core/repository"
	"context"
)

type (
	SaveData interface {
		Execute(ctx *context.Context, p *domain.Pedido) (*domain.Pedido, error)
	}
	saveData struct{
		repo repository.PedidoRepository
	}
)

func NewSaveData(repo repository.PedidoRepository) SaveData{
	return &saveData{repo: repo}
}

func (s *saveData) Execute(ctx *context.Context, p *domain.Pedido) (*domain.Pedido, error){
	out, err:= s.repo.Add(ctx, p)
	if err != nil{
		return &domain.Pedido{}, err
	}
	return &out, nil
}