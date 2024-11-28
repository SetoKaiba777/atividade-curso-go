package usecase

import (
	"api-pedidos/core/domain"
	"api-pedidos/core/usecase/input"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFindByIdExecute(t *testing.T) {
	respDef := domain.Pedido{
		Id:     "1",
		UserId: "testUser",
		ListaProdutos: []domain.Produto{
			{
				Nome:       "Bomba Nuclear",
				Valor:      10.00,
				Quantidade: 10,
			},
		},
	}

	tt := []struct {
		name     string
		input    *input.FindByIdInput
		expected domain.Pedido
		err      error
	}{
		{
			name:     "Sucesso",
			input:    &input.FindByIdInput{Id: "1"},
			expected: respDef,
			err:      nil,
		},
		{
			name:     "Erro",
			input:    &input.FindByIdInput{Id: "1"},
			expected: domain.Pedido{},
			err:      databaseErr,
		},
	}
	for _, sc := range tt {
		t.Run(sc.name, func(t *testing.T) {
			//Cria a estrtura de teste
			r := NewMockDatabase()
			uc := NewSearchById(r)

			// Estabelece comportamento do mock
			r.On("GetById", mock.Anything, mock.Anything).Return(sc.expected, sc.err)

			//Execução
			ctx := context.TODO()
			out, err := uc.Execute(&ctx, sc.input)

			//Validação
			assert.Equal(t, sc.expected, *out)
			assert.Equal(t, sc.err, err)
		})
	}
}
