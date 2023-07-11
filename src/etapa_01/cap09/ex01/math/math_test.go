package math

import "testing"

type testpair struct {
	values  []float64
	average float64
}

var tests = []testpair{
	{[]float64{1, 1, 1}, 1},
	{[]float64{1, 2}, 1.5},
	{[]float64{}, 0},
}

func TestAverage(t *testing.T) {
	for _, v := range tests {
		res := Average(v.values)
		if res != v.average {
			t.Error("Para: ", v.values,
				"Esperado: ", v.average,
				"Obtido: ", res)
		}
	}
}
