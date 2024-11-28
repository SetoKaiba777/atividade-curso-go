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

func TestUpdateStatusExecute(t *testing.T) {
	expected := domain.Pedido{
		Id:         "1",
		UserId:     "testUser",
		Status:     domain.PROCESSADO,
		ValorTotal: 100.00,
		ListaProdutos: []domain.Produto{
			{
				Nome:       "Bomba Nuclear",
				Valor:      10.00,
				Quantidade: 10,
			},
		},
	}

	outPadrao := domain.Pedido{
		Id:         "1",
		UserId:     "testUser",
		Status:     domain.PENDENTE,
		ValorTotal: 100.00,
		ListaProdutos: []domain.Produto{
			{
				Nome:       "Bomba Nuclear",
				Valor:      10.00,
				Quantidade: 10,
			},
		},
	}

	tt := []struct {
		name        string
		input       *input.UpdateStatusInput
		expected    domain.Pedido
		expectedErr error
		addMock     domain.Pedido
		addErr      error
		getMock     domain.Pedido
		getErr      error
	}{
		{
			name:        "Sucesso",
			input:       &input.UpdateStatusInput{Id: "1", Status: domain.PROCESSADO},
			expected:    expected,
			expectedErr: nil,
			getMock:     outPadrao,
			getErr:      nil,
			addMock:     expected,
			addErr:      nil,
		},
		{
			name:        "Falha buscar dados",
			input:       &input.UpdateStatusInput{Id: "1", Status: domain.PROCESSADO},
			expected:    domain.Pedido{},
			expectedErr: databaseErr,
			getMock:     outPadrao,
			getErr:      databaseErr,
		},
		{
			name:        "Falha ao salvar dados",
			input:       &input.UpdateStatusInput{Id: "1", Status: domain.PROCESSADO},
			expected:    domain.Pedido{},
			expectedErr: databaseErr,
			getMock:     outPadrao,
			getErr:      nil,
			addMock:     expected,
			addErr:      databaseErr,
		},
		{
			name:        "Falha ao atualizar status",
			input:       &input.UpdateStatusInput{Id: "1", Status: domain.CONCLUIDO},
			expected:    domain.Pedido{},
			expectedErr: fmt.Errorf("mudança de status inválida! %v", domain.CONCLUIDO),
			getMock:     outPadrao,
			getErr:      nil,
		},
	}
	for _, sc := range tt {
		t.Run(sc.name, func(t *testing.T) {
			//Cria a estrtura de teste
			r := NewMockDatabase()
			uc := NewUpdateStatus(r)

			// Estabelece comportamento do mock
			r.On("GetById", mock.Anything, mock.Anything).Return(sc.getMock, sc.getErr)
			r.On("Add", mock.Anything, mock.Anything).Return(sc.addMock, sc.addErr)

			//Execução
			ctx := context.TODO()
			out, err := uc.Execute(&ctx, sc.input)

			//Validação
			assert.Equal(t, sc.expected, *out)
			assert.Equal(t, sc.expectedErr, err)
		})
	}
}
