package controller

import (
	"api-pedidos/core/domain"
	"api-pedidos/core/erros"
	"api-pedidos/core/usecase/input"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateStatusController(t *testing.T) {
	uc := updateStatusMock{}

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
			requestBody: input.UpdateStatusInput{Id: "10", Status: 3},
			mockUseCaseSetup: func() {
				uc.On("Execute", mock.Anything, mock.Anything).Return(&domain.Pedido{}, erros.NewChangeStatusErr(3)).Once()
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:        "success",
			requestBody: input.UpdateStatusInput{Id: "10", Status: 3},
			mockUseCaseSetup: func() {
				uc.On("Execute", mock.Anything, mock.Anything).Return(&domain.Pedido{Id: "10", UserId: "101", Status: domain.Status(3)}, nil).Once()
			},
			expectedStatusCode: http.StatusOK,
		},
	}
	for _, sc := range tt {
		t.Run(sc.name, func(t *testing.T) {
			c := NewUpdateStatusController(&uc)

			var body io.Reader
			if sc.requestBody != nil {
				jsonData, _ := json.Marshal(sc.requestBody)
				body = bytes.NewBuffer(jsonData)
			} else {
				body = &errorReader{}
			}
			req := httptest.NewRequest(http.MethodPut, "/v1/pedidos", body)
			w := httptest.NewRecorder()

			sc.mockUseCaseSetup()

			c.Execute(w, req)
			assert.Equal(t, sc.expectedStatusCode, w.Code)
		})
	}
}
