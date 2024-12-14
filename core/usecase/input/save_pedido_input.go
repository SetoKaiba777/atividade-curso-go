package input

import "api-pedidos/core/domain"

type SaveInput struct {
	UserId        string
	Status        domain.Status
	ListaProdutos []domain.Produto
}
