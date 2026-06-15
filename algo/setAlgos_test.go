package algo_test

import (
	"fmt"
	"testing"

	"github.com/Shangin-Leonid/acrogen/algo"
	"github.com/stretchr/testify/assert"
)

func ExampleCalcOrderedCartesianProduct() {

	upperCase := []string{"A", "B"}
	lowerCase := []string{"a", "b"}
	letterProd := algo.CalcOrderedCartesianProduct([][]string{upperCase, lowerCase})
	fmt.Println(letterProd)

	familyOfSets := [][]int{
		[]int{1, 2}, []int{3, 4}, []int{5, 6},
	}
	numProd := algo.CalcOrderedCartesianProduct(familyOfSets)
	fmt.Println(numProd)

	// Output:
	// [[A a] [A b] [B a] [B b]]
	// [[1 3 5] [1 3 6] [1 4 5] [1 4 6] [2 3 5] [2 3 6] [2 4 5] [2 4 6]]
}

func Test_CalcOrderedCartesianProduct(t *testing.T) {
	testSuites := []struct {
		name            string
		setFamilyParam  [][]int
		productExpected [][]int
	}{
		{"empty family", [][]int{}, nil},

		{"one set",
			[][]int{
				[]int{1, 2, 3},
			},
			[][]int{
				[]int{1}, []int{2}, []int{3},
			},
		},

		{"one element set",
			[][]int{
				[]int{1}, []int{2, 3}, []int{4, 5},
			},
			[][]int{
				[]int{1, 2, 4}, []int{1, 2, 5}, []int{1, 3, 4}, []int{1, 3, 5},
			},
		},

		{"one of sets is empty",
			[][]int{
				[]int{1, 2, 3}, []int{}, []int{7, 8, 9},
			},
			[][]int{},
		},

		{"two sets",
			[][]int{
				[]int{1, 2}, []int{8, 9},
			},
			[][]int{
				[]int{1, 8}, []int{1, 9}, []int{2, 8}, []int{2, 9},
			},
		},

		{"three sets",
			[][]int{
				[]int{1, 2}, []int{3, 4}, []int{5, 6},
			},
			[][]int{
				[]int{1, 3, 5}, []int{1, 3, 6}, []int{1, 4, 5}, []int{1, 4, 6}, []int{2, 3, 5}, []int{2, 3, 6}, []int{2, 4, 5}, []int{2, 4, 6},
			},
		},
	}

	for _, ts := range testSuites {
		t.Run(ts.name, func(t *testing.T) {

			// Test code
			res := algo.CalcOrderedCartesianProduct(ts.setFamilyParam)
			assert.Equal(t, res, ts.productExpected)

		})
	}
}
