package controller

import (
	"api-pedidos/core/domain"
	"api-pedidos/core/usecase/input"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSaveController(t *testing.T) {

	var uc saveControllerMock

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
		name               string
		requestBody        any
		mockUseCaseSetup   func()
		expectedStatusCode int
	}{
		{
			name:               "body read error",
			requestBody:        nil,
			mockUseCaseSetup:   func() {},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name:               "Invalid Json",
			requestBody:        "{teste: 01",
			mockUseCaseSetup:   func() {},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name:        "use case Error",
			requestBody: inPadrao,
			mockUseCaseSetup: func() {
				uc.On("Execute", mock.Anything, mock.Anything).Return(&domain.Pedido{}, errors.New("database error")).Once()
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name:        "success",
			requestBody: inPadrao,
			mockUseCaseSetup: func() {
				uc.On("Execute", mock.Anything, mock.Anything).Return(&outPadrao, nil).Once()
			},
			expectedStatusCode: http.StatusOK,
		},
	}
	for _, sc := range tt {
		t.Run(sc.name, func(t *testing.T) {
			c := NewSaveController(&uc)

			var body io.Reader
			if sc.requestBody != nil {
				jsonData, _ := json.Marshal(sc.requestBody)
				body = bytes.NewBuffer(jsonData)
			} else {
				body = &errorReader{}
			}
			req := httptest.NewRequest(http.MethodPost, "/v1/pedidos", body)
			w := httptest.NewRecorder()

			sc.mockUseCaseSetup()

			c.Execute(w, req)
			assert.Equal(t, sc.expectedStatusCode, w.Code)
		})
	}
}
