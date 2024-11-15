package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNext(t *testing.T) {

	tt := []struct{
			name string
			input int
			expected int
		}{

			{
				name: "Sucesso",
				input: 0,
				expected: 1,
			},
			{
				name: "Falha",
				input: 3,
				expected: 3,
			},
		}
		for _, scenario := range tt {
			t.Run(scenario.name, func(t *testing.T) {
				out := Status(scenario.input).Next()
                ex := Status(scenario.expected)		
				
 				assert.Equal(t,ex, out)
			})
		}
}

func TestPrevious(t *testing.T) {
	tt := []struct{
		name string
		input int
		expected int
	}{

		{
			name: "Sucesso",
			input: 3,
			expected: 2,
		},
		{
			name: "Falha",
			input: 0,
			expected: 0,
		},
	}
	for _, scenario := range tt {
		t.Run(scenario.name, func(t *testing.T) {
			out := Status(scenario.input).Previous()
			ex := Status(scenario.expected)		
			
			 assert.Equal(t,ex, out)
		})
	}
}
