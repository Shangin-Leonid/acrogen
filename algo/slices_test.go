package algo_test

import (
	"fmt"

	"github.com/Shangin-Leonid/acrogen/algo"
)

func ExampleGetCopy() {

	letters := []string{"A", "B", "C", "X", "Y", "Z"}
	lettersCopy := algo.GetCopy(letters)
	fmt.Println(lettersCopy)

	// Output:
	// [A B C X Y Z]
}

func ExampleReverseSlice() {

	letters := []string{"A", "B", "C", "X", "Y", "Z"}
	algo.ReverseSlice(letters)
	fmt.Println(letters)

	// Output:
	// [Z Y X C B A]
}

func ExampleSplitSlice() {

	res := algo.SplitSlice(make([]int, 5), -8)
	fmt.Println(res) // []

	res = algo.SplitSlice(make([]int, 0), 0)
	fmt.Println(res) // []

	res = algo.SplitSlice(make([]int, 0), 2)
	fmt.Println(res) // []

	res = algo.SplitSlice(make([]int, 1), 0)
	fmt.Println(res) // []

	res = algo.SplitSlice(make([]int, 1), 1)
	fmt.Println(res) // [{0 1}]

	res = algo.SplitSlice(make([]int, 1), 2)
	fmt.Println(res) // [{0 1}]

	res = algo.SplitSlice(make([]int, 10), 5)
	fmt.Println(res) // [{0 2} {2 4} {4 6} {6 8} {8 10}]

	res = algo.SplitSlice(make([]int, 10), 4)
	fmt.Println(res) // [{0 3} {3 6} {6 8} {8 10}]

	res = algo.SplitSlice(make([]int, 3), 100)
	fmt.Println(res) // [{0 1} {1 2} {2 3}]

	// Output:
	// []
	// []
	// []
	// []
	// [{0 1}]
	// [{0 1}]
	// [{0 2} {2 4} {4 6} {6 8} {8 10}]
	// [{0 3} {3 6} {6 8} {8 10}]
	// [{0 1} {1 2} {2 3}]
}
