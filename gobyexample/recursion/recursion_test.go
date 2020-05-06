package recursion

import "testing"

func TestFact(t *testing.T) {
	var tests = []struct {
		input int
		want  int
	}{
		{1, 1},
		{2, 2},
		{3, 6},
		//{3, 5},
	}

	for _, test := range tests {
		if got := fact(test.input); got != test.want {
			t.Errorf("fact(%d) == %d, want %d", test.input, got, test.want)
		}
	}
}

//
func BenchmarkFact(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fact(10)
	}
}
