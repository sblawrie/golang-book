package math

import "testing"

type testpair struct {
    values []float64
    result float64
}

var tests = []testpair{
    { []float64{1,2}, 1.5 },
    { []float64{1,1,1,1,1,1}, 1 },
    { []float64{-1,1}, 0 },
}

var min_tests = []testpair{
	{ []float64{1,2}, 1 },
	{ []float64{1,1,1,1,1,1}, 1 },
	{ []float64{-1,1}, -1 },
}

var max_tests = []testpair{
	{ []float64{1,2}, 2 },
	{ []float64{1,1,1,1,1,1}, 1 },
	{ []float64{-1,1}, 1 },
}

func TestAverage(t *testing.T) {
    for _, pair := range tests {
        v := Average(pair.values)
        if v != pair.result {
            t.Error(
                "For", pair.values, 
                "expected", pair.result,
                "got", v,
            )
        }
    }
}

func TestMin(t *testing.t) {
	for _, pair := range min_tests {
		v := Min(pair.values)
		if v != pair.result {
			t.Error(
				"For", pair.values,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}

func TestMax(t *testing.t) {
	for _, pair := range max_tests {
		v := Max(pair.values)
		if v != pair.result {
			t.Error(
				"For", pair.values,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}
