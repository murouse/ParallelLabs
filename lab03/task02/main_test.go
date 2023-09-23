package main

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCorrCoef(t *testing.T) {
	testTable := []struct {
		name string
		x    []float64
		y    []float64
		exp  float64
	}{
		{
			name: "1",
			x:    []float64{1, 4, 2, 4, 5, 8, 7, 4, 7, 9, 1, 3},
			y:    []float64{2, 4, 5, 8, 7, 4, 7, 9, 1, 3, 5, 6},
			exp:  -0.14786316657163495,
		},
		{
			name: "2",
			x:    []float64{1, 4, 2, 4, 5, 8},
			y:    []float64{7, 9, 1, 3, 5, 6},
			exp:  0.17142857142857143,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			act := corrCoef(test.x, test.y)
			require.Equal(t, test.exp, act)
		})
	}
}

func TestAcf(t *testing.T) {
	x := []float64{1, 4, 2, 4, 5, 8, 7, 4, 7, 9, 1, 3, 5, 6}
	var res []float64
	const lags = 9

	for lag := 0; lag <= lags; lag++ {
		l, r := x[:len(x)-lag], x[lag:]
		fmt.Println(l, r)
		coef := corrCoef(l, r)
		fmt.Println(coef)
		res = append(res, coef)
	}

	fmt.Println(res)

	//x := []uint64{1, 4, 2, 4, 5, 8, 7, 4, 7, 9, 1, 3, 5, 6}
	//plotAcf(x, 9)
}
