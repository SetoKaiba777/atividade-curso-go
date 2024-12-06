package erros

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidRequestErr(t *testing.T) {
	test := struct {
		name string
		err  string
	}{
		name: "Testando funçao de erro",
		err:  "requisicao inválida!",
	}
	t.Run(test.name, func(t *testing.T) {
		err := NewInvalidRequestErr()
		assert.Equal(t, test.err, err.Error())
	})
}
