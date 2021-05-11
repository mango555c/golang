package Iterations

import "testing"

func TestIteration(t *testing.T) {
	letters := Iteration("a")
	want := "aaaaa"

	if letters != want {
		t.Errorf("want '%q' but got '%q'", letters, want)
	}
}

func BenchmarkIteration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Iteration("a")
	}
}
