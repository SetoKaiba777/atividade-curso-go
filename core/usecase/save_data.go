package usecase

import (
	"api-pedidos/core/domain"
	"api-pedidos/core/repository"
	"api-pedidos/core/usecase/input"
	"context"
)

type (
	SaveData interface {
		Execute(ctx *context.Context, p *input.SaveInput) (*domain.Pedido, error)
	}
	saveData struct {
		repo repository.PedidoRepository
	}
)

func NewSaveData(repo repository.PedidoRepository) SaveData {
	return &saveData{repo: repo}
}

func (s *saveData) Execute(ctx *context.Context, i *input.SaveInput) (*domain.Pedido, error) {
	p := domain.NewPedido(
		domain.WithUserId(i.UserId),
		domain.WithListaProdutos(i.ListaProdutos),
		domain.WithStaus(i.Status),
	)

	out, err := s.repo.Add(ctx, p)
	if err != nil {
		return &domain.Pedido{}, err
	}
	return &out, nil
}
