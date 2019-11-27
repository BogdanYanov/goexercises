package fibonacci

import "testing"

type testpair struct {
	num uint64
	res uint64
}

var testPairs = []testpair{
	{0, 0},
	{1, 1},
	{2, 1},
	{3, 2},
	{5, 5},
	{6, 8},
	{9, 34},
	{14,377},
	{15,610},
	{30,832040},
}

func TestFibonacci(t *testing.T) {
	for _, pair := range testPairs {
		val := Fibonacci(pair.num)

		if val != pair.res {
			t.Error(
				"For", pair.num,
				"expected", pair.res,
				"got", val)
		}
	}
}
