package controller

import (
	"api-pedidos/core/domain"
	"api-pedidos/core/erros"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSearchByOrderIdController(t *testing.T) {
	uc := searchOrderByIdMock{}

	tt := []struct {
		name               string
		input              string
		mockUseCaseSetup   func()
		expectedStatusCode int
	}{
		{
			name:               "erro de url",
			input:              "/v1/usuario?orderId=",
			mockUseCaseSetup:   func() {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:  "erro not found",
			input: "/v1/usuario?orderId=10",
			mockUseCaseSetup: func() {
				uc.On("Execute", mock.Anything, mock.Anything).Return(&domain.Pedido{}, erros.NewNotFoundErr("10", "Usuário")).Once()
			},
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:  "sucesso",
			input: "/v1/usuario?orderId=1",
			mockUseCaseSetup: func() {
				uc.On("Execute", mock.Anything, mock.Anything).Return(&domain.Pedido{
					Id:     "1",
					UserId: "testUser",
					ListaProdutos: []domain.Produto{
						{
							Nome:       "Bomba Nuclear",
							Valor:      10.00,
							Quantidade: 10,
						},
					}}, nil).Once()
			},
			expectedStatusCode: http.StatusOK,
		},
	}
	for _, sc := range tt {
		t.Run(sc.name, func(t *testing.T) {
			c := NewSearchByOrderIdController(&uc)

			r := httptest.NewRequest("GET", sc.input, &bytes.Reader{})
			w := httptest.NewRecorder()

			sc.mockUseCaseSetup()
			c.Execute(w, r)

			assert.Equal(t, sc.expectedStatusCode, w.Code)
		})
	}
}
