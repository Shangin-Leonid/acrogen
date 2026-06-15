package algo_test

import (
	"fmt"

	"github.com/Shangin-Leonid/acrogen/algo"
)

func ExampleCalcFactorial() {
	fmt.Println(algo.CalcFactorial(0))
	fmt.Println(algo.CalcFactorial(1))
	fmt.Println(algo.CalcFactorial(2))
	fmt.Println(algo.CalcFactorial(3))
	fmt.Println(algo.CalcFactorial(10))

	// Output:
	// 1
	// 1
	// 2
	// 6
	// 3628800
}
