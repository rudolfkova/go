package quests

import "testing"

type testpair struct {
	value  int
	result int
}

var tests = []testpair{
	{1, 1},
	{0, 0},
	{2, 2},
}

func TestFib(t *testing.T) {
	for _, pair := range tests {
		v := Fib(pair.value)
		if v != pair.result {
			t.Error(
				"For", pair.value,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}
