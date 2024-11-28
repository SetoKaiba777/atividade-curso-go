package domain

import (
	"api-pedidos/core/erros"
)

type (
	Pedido struct {
		Id            string
		UserId        string
		Status        Status
		ListaProdutos []Produto
		ValorTotal    float64
	}
	Produto struct {
		Nome       string
		Valor      float64
		Quantidade int
	}
	Opt func(*Pedido)
)

func NewPedido(opts ...Opt) *Pedido {
	p := &Pedido{ListaProdutos: []Produto{}}
	for _, opt := range opts {
		opt(p)
	}
	p.valorTotal()

	return p
}

func WithUserId(userId string) Opt {
	return func(p *Pedido) {
		p.UserId = userId
	}
}

func WithStaus(status Status) Opt {
	return func(p *Pedido) {
		p.Status = status
	}
}

func WithListaProdutos(produtos []Produto) Opt {
	return func(p *Pedido) {
		p.ListaProdutos = produtos
	}
}

func (p *Pedido) valorTotal() {
	total := 0.0
	for _, v := range p.ListaProdutos {
		total += float64(v.Quantidade) * v.Valor
	}
	p.ValorTotal = total
}

func (p *Pedido) UpdateStatus(status Status) error {
	if status == p.Status.Next() || status == p.Status.Previous() {
		p.Status = status
		return nil
	}
	return erros.NewChangeStatusErr(int(status))
}
