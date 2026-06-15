package algo_test

import (
	"testing"

	"github.com/Shangin-Leonid/acrogen/algo"
)

func Test_CalcFactorial(t *testing.T) {
	testSuites := []struct {
		name     string
		nParam   uint
		expected uint
	}{
		{"zero", 0, 1},
		{"one", 1, 1},
		{"two", 2, 2},
		{"three", 3, 6},
		{"ten", 10, 3628800},
	}

	for _, ts := range testSuites {
		t.Run(ts.name, func(t *testing.T) {

			// Test code
			res := algo.CalcFactorial(ts.nParam)
			if res != ts.expected {
				t.Errorf("Expected: %d, got: %d", ts.expected, res)
			}

		})
	}
}
