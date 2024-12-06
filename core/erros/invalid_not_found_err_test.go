package erros

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotFoundErr(t *testing.T) {
	test := struct {
		name string
		err string
	}{
		name: "Testando funçao de erro",
		err: fmt.Sprintf("%s não encontrado, id : %s","Usuário", "10"),
	}
	t.Run(test.name, func(t *testing.T) {
		err := NewNotFoundErr("Usuário","10")
		assert.Equal(t, test.err, err.Error())
	})
}