package math

import "testing"

type testpair struct {
	values  []float64
	average float64
}

var tests = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1, 1}, 1},
	{[]float64{1, 1, 1, 1, 1, 1, 10}, 1},
	{[]float64{-1, 1}, 0},
}

func TestAverage(t *testing.T) {

	for _, pair := range tests {
		v := Average(pair.values)
		if v != pair.average {
			t.Error("Para: ", pair.values,
				"Esperado: ", pair.average,
				"Obtido: ", v)
		}
	}
}
