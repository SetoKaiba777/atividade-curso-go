package usecase

import (
	"api-pedidos/core/domain"
	"api-pedidos/core/usecase/input"
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSaveDataExecute(t *testing.T) {
	inPadrao := input.SaveInput{
		UserId: "testUser",
		ListaProdutos: []domain.Produto{
			{
				Nome:       "Bomba Nuclear",
				Valor:      10.00,
				Quantidade: 10,
			},
		},
	}

	outPadrao := domain.Pedido{
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
		input    input.SaveInput
		expected domain.Pedido
		err      error
	}{
		{
			name:     "Sucesso",
			input:    inPadrao,
			expected: outPadrao,
			err:      nil,
		},
		{
			name:     "Erro",
			input:    inPadrao,
			expected: domain.Pedido{},
			err:      fmt.Errorf("Database error"),
		},
	}
	for _, sc := range tt {
		t.Run(sc.name, func(t *testing.T) {
			//Cria a estrtura de teste
			r := NewMockDatabase()
			uc := NewSaveData(r)

			// Estabelece comportamento do mock
			r.On("Add", mock.Anything, mock.Anything).Return(sc.expected, sc.err)

			//Execução
			ctx := context.TODO()
			out, err := uc.Execute(&ctx, &sc.input)

			//Validação
			assert.Equal(t, sc.expected, *out)
			assert.Equal(t, sc.err, err)
		})
	}
}
