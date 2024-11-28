package usecase

import (
	"api-pedidos/core/domain"
	"api-pedidos/core/repository"
	"api-pedidos/core/usecase/input"
	"context"
)

type (
	UpdateStatus interface {
		Execute(ctx *context.Context, i *input.UpdateStatusInput) (*domain.Pedido, error)
	}
	updateStatus struct {
		repo repository.PedidoRepository
	}
)

func NewUpdateStatus(repo repository.PedidoRepository) UpdateStatus {
	return &updateStatus{repo: repo}
}

func (s *updateStatus) Execute(ctx *context.Context, i *input.UpdateStatusInput) (*domain.Pedido, error) {
	out, err := s.repo.GetById(ctx, i.Id)
	if err != nil {
		return &domain.Pedido{}, err
	}

	err = out.UpdateStatus(i.Status)
	if err != nil {
		return &domain.Pedido{}, err
	}
	_, err = s.repo.Add(ctx, &out)
	if err != nil {
		return &domain.Pedido{}, err
	}
	return &out, nil
}
