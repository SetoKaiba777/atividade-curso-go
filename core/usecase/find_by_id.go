package usecase

import (
	"api-pedidos/core/domain"
	"api-pedidos/core/repository"
	"api-pedidos/core/usecase/input"
	"context"
)

type (
	SearchById interface {
		Execute(ctx *context.Context, i *input.FindByIdInput) (*domain.Pedido, error)
	}
	searchById struct {
		repo repository.PedidoRepository
	}
)

func NewSearchById(repo repository.PedidoRepository) SearchById {
	return &searchById{repo: repo}
}

func (s *searchById) Execute(ctx *context.Context, i *input.FindByIdInput) (*domain.Pedido, error) {
	out, err := s.repo.GetById(ctx, i.Id)
	if err != nil {
		return &domain.Pedido{}, err
	}
	return &out, nil
}
