package sts

import "testing"

type testPair struct {
	values []float64
	sum    float64
	avg    float64
	min    float64
	median float64
	max    float64
}

var tests = []testPair{
	{[]float64{1.}, 1., 1., 1., 1., 1.},
	{[]float64{1., 2.}, 3., 1.5, 1., 1.5, 2.},
	{[]float64{-1., 1.}, 0., 0., -1, 0, 1},
	{[]float64{1., 2., 3.}, 6., 2., 1, 2, 3},
}

func testSomaInt(t *testing.T) {
	for _, pair := range tests {
		total := SomaFloat64(pair.values)
		if total != pair.sum {
			t.Error(
				"Para:", pair.values,
				"Experado:", pair.sum,
				"Realizado:", total,
			)
		}
	}
}

func testMediaFloat(t *testing.T) {
	for _, pair := range tests {
		total := AvgFloat64(pair.values)
		if total != pair.avg {
			t.Error(
				"Para:", pair.values,
				"Experado:", pair.avg,
				"Realizado:", total,
			)
		}
	}
}

func testMinFloat(t *testing.T) {
	for _, pair := range tests {
		total := MinFloat64(pair.values)
		if total != pair.min {
			t.Error(
				"Para:", pair.values,
				"Experado:", pair.min,
				"Realizado:", total,
			)
		}
	}
}

func testMaxFloat(t *testing.T) {
	for _, pair := range tests {
		total := MaxFloat64(pair.values)
		if total != pair.max {
			t.Error(
				"Para:", pair.values,
				"Experado:", pair.max,
				"Realizado:", total,
			)
		}
	}
}

func testMedianFloat(t *testing.T) {
	for _, pair := range tests {
		total := MedianFloat64(pair.values)
		if total != pair.median {
			t.Error(
				"Para:", pair.values,
				"Experado:", pair.median,
				"Realizado:", total,
			)
		}
	}
}
