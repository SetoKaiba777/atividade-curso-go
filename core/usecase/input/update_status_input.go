package input

import "api-pedidos/core/domain"

type UpdateStatusInput struct {
	Id     string        `json:"id"`
	Status domain.Status `json:"status"`
}
