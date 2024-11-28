package erros

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChangeStatusErr(t *testing.T) {
	test := struct {
		name string
		status int
		err string
	}{
		name: "Testando funçao de erro",
		status: 1,
		err: fmt.Sprintf("mudança de status inválida! %v", 1),
	}
	t.Run(test.name, func(t *testing.T) {
		err := NewChangeStatusErr(test.status)
		assert.Equal(t, test.err, err.Error())
	})
}