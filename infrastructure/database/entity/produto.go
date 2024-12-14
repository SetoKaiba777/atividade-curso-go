package entity

import (
	"api-pedidos/core/domain"
	"encoding/json"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	Pedido struct {
		Id            string  `gorm:"type:char(36);primaryKey"`
		UserId        string  `gorm:"type:char(36);not null;index"`
		Status        int     `gorm:"type:int;not null"`
		ListaProdutos string  `gorm:"type:text;not null"`
		ValorTotal    float64 `gorm:"type:double;not null"`
	}
)

func (p *Pedido) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.New().String()
	p.Id = uuid
	return nil
}

func NewPedidoEntity(p *domain.Pedido) (*Pedido, error) {
	pList, err := json.Marshal(p.ListaProdutos)
	if err != nil {
		return &Pedido{}, err
	}
	return &Pedido{UserId: p.UserId, Status: int(p.Status), ValorTotal: p.ValorTotal, ListaProdutos: string(pList)}, nil
}

func (e Pedido) ToDomain() (*domain.Pedido, error) {
	var lista []domain.Produto
	if err := json.Unmarshal([]byte(e.ListaProdutos), &lista); err != nil {
		return &domain.Pedido{}, err
	}
	return &domain.Pedido{Id: e.Id, UserId: e.UserId, Status: domain.Status(e.Status), ValorTotal: e.ValorTotal, ListaProdutos: lista}, nil
}
