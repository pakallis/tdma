package tdma_test

import (
	"testing"
	"github.com/go-test/deep"
	. "github.com/pakallis/mathutils/tdma"
)

func TestSolve(t *testing.T) {
	var l = []float64{0, -1, -1, -1}
	var d = []float64{4, 4, 4, 4}
	var u = []float64{-1, -1, -1, 0}
	var b = []float64{5, 5, 10, 23}
	var expected = []float64{2, 3, 5, 7}
	result := Solve(l, d, u, b)
	if diff := deep.Equal(result, expected); diff != nil {
		t.Error(diff)
	}
}


func BenchmarkSolve(bn *testing.B) {
	var l = []float64{0, -1, -1, -1}
	var d = []float64{4, 4, 4, 4}
	var u = []float64{-1, -1, -1, 0}
	var b = []float64{5, 5, 10, 23}
	for i := 0; i < bn.N; i++ {
		Solve(l, d, u, b)
	}
}
