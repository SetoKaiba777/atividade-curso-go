package domain

import (
	"api-pedidos/core/erros"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPedido(t *testing.T) {
	lp := []Produto{{"Meu Produto", 10.00, 10}}
	s := PROCESSADO
	uId := "Meu Id"

	pedido := NewPedido(
		WithStaus(s),
		WithUserId(uId),
		WithListaProdutos(lp),
	)

	assert.Equal(t, pedido.UserId, uId)
	assert.Equal(t, pedido.Id, "")
	assert.Equal(t, pedido.Status, s)
	assert.Equal(t, pedido.ListaProdutos, lp)
	assert.Equal(t, pedido.ValorTotal, 100.00)
}

func TestUpdateStatus(t *testing.T) {
	tt := []struct {
		name   string
		status Status
		err    error
	}{
		{
			name:   "Valid transtion",
			status: PROCESSADO,
			err:    nil,
		},
		{
			name:   "Invalid transtion",
			status: ENVIADO,
			err:    erros.NewChangeStatusErr(int(ENVIADO)),
		},
	}
	for _, sc := range tt {
		t.Run(sc.name, func(t *testing.T) {
			p := NewPedido(
				WithStaus(PENDENTE),
			)

			err := p.UpdateStatus(sc.status)
			if err != nil {
				assert.Equal(t, sc.err.Error(), err.Error())
			}
		})
	}
}
