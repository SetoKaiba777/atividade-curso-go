package erros

import (
	"fmt"
)

type InvalidRequestErr struct {

}

func NewInvalidRequestErr() error {
	return &ChangeStatusErr{}
}

func (c *InvalidRequestErr) Error() string {
	return fmt.Sprintf("requisição inválida! ")
}
