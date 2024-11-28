package erros

import (
	"fmt"
)

type ChangeStatusErr struct {
	status int
}

func NewChangeStatusErr(status int) error {
	return ChangeStatusErr{status: status}
}

func (c ChangeStatusErr) Error() string {
	return fmt.Sprintf("mudança de status inválida! %v", c.status)
}
