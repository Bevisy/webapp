package myatoi

import (
	"testing"
)

func TestMyAtoi(t *testing.T) {
	var tests = []struct {
		input string
		ans   int
	}{
		{"123", 123},
		{"321", 321},
	}

	for _, test := range tests {
		got := myatoi(test.input)
		if got != test.ans {
			t.Errorf("input is %s, expect %d, got %d\n", test.input,
				test.ans, got)
		}
	}
}
