package usecase

import (
	"api-pedidos/core/domain"
	"api-pedidos/core/repository"
	"api-pedidos/core/usecase/input"
	"context"
)

type (
	SearchByUserId interface {
		Execute(ctx *context.Context, i *input.FindByIdInput) (*[]domain.Pedido, error)
	}
	searchByUserId struct {
		repo repository.PedidoRepository
	}
)

func NewSearchByUserId(repo repository.PedidoRepository) SearchByUserId {
	return &searchByUserId{repo: repo}
}

func (s *searchByUserId) Execute(ctx *context.Context, i *input.FindByIdInput) (*[]domain.Pedido, error) {
	out, err := s.repo.GetByUserId(ctx, i.Id)
	if err != nil {
		return &[]domain.Pedido{}, err
	}
	return &out, nil
}
