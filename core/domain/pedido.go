package domain

import "fmt"

type (
	Pedido struct {
		Id            string
		UserId        string
		Status        Status
		ListaProdutos []Produto
		ValorTotal int64
	}
	Produto struct {
		Nome       string
		Valor      float64
		Quantidade int
	}
)
func NewPedido()

func (p *Pedido) valorTotal() float64 {
	total := 0.0
	for _, v := range p.ListaProdutos {
		total += float64(v.Quantidade) * v.Valor
	}
	return total
}

func (p *Pedido) UpdateStatus(status Status) error {
	if status == p.Status.Next() || status == p.Status.Previous() {
		p.Status = status
		return nil
	}
	return fmt.Errorf("mudança de status inválida! %v", status)
}