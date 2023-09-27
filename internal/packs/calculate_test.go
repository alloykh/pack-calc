package packs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPacks(t *testing.T) {

	packsCalc := NewPacksCalculator(Config{Packs: []int{
		250,
		500,
		1000,
		2000,
		5000,
	}})

	for _, c := range []struct {
		name          string
		input         int
		expectedM     map[int]int
		expectedTotal int
		err           error
	}{
		{
			name:  "1",
			input: 0,
			err:   errInputMustBeGreaterThanZero,
		},
		{
			name:  "1",
			input: 1,
			expectedM: map[int]int{
				250: 1,
			},
			expectedTotal: 1,
		},

		{
			name:  "250",
			input: 250,
			expectedM: map[int]int{
				250: 1,
			},
			expectedTotal: 1,
		},

		{
			name:  "251",
			input: 251,
			expectedM: map[int]int{
				500: 1,
			},
			expectedTotal: 1,
		},

		{
			name:  "500",
			input: 500,
			expectedM: map[int]int{
				500: 1,
			},
			expectedTotal: 1,
		},

		{
			name:  "501",
			input: 501,
			expectedM: map[int]int{
				500: 1,
				250: 1,
			},
			expectedTotal: 2,
		},

		{
			name:  "12001",
			input: 12001,
			expectedM: map[int]int{
				5000: 2,
				2000: 1,
				250:  1,
			},
			expectedTotal: 4,
		},
		{
			name:  "1000",
			input: 1000,
			expectedM: map[int]int{
				1000: 1,
			},
			expectedTotal: 1,
		},
	} {
		t.Run(c.name, func(t *testing.T) {

			actM, actTotal, err := packsCalc.Calculate(c.input)

			require.Equal(t, c.err, err)
			require.Equal(t, c.expectedM, actM)
			require.Equal(t, c.expectedTotal, actTotal)
		})
	}

}
