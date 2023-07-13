package math

import "testing"

type testpair struct {
	values  []float64
	average float64
	min     float64
	max     float64
}

var tests = []testpair{
	{[]float64{1, 1, 1}, 1, 1, 1},
	{[]float64{1, 2}, 1.5, 1, 2},
	{[]float64{-10, 10}, 0, -10, 10},
	{[]float64{}, 0, 0, 0},
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

func TestMin(t *testing.T) {
	for _, v := range tests {
		res := Min(v.values)
		if res != v.min {
			t.Error("Para: ", v.values,
				"Esperado: ", v.min,
				"Obtido: ", res)
		}
	}
}

func TestMax(t *testing.T) {
	for _, v := range tests {
		res := Max(v.values)
		if res != v.max {
			t.Error("Para: ", v.values,
				"Esperado: ", v.max,
				"Obtido: ", res)
		}
	}
}
