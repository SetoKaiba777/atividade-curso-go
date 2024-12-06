package erros

import "fmt"

type NotFoundErr struct {
	id   string
	tipo string
}

func NewNotFoundErr(tipo, id string) error {
	return NotFoundErr{
		id:   id,
		tipo: tipo,
	}
}

func (e NotFoundErr) Error() string {
	return fmt.Sprintf("%s não encontrado, id : %s", e.tipo, e.id)
}
