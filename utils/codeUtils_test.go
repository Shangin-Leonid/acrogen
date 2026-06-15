package utils_test

import (
	"fmt"

	"github.com/Shangin-Leonid/acrogen/utils"
)

func ExampleTerOp() {

	fmt.Println(utils.TerOp(true, 1, 0))
	fmt.Println(utils.TerOp(2*2 == 5, "A", "B"))

	// Output:
	// 1
	// B
}

func ExampleAbsInt() {

	fmt.Println(utils.AbsInt(0))
	fmt.Println(utils.AbsInt(-1))
	fmt.Println(utils.AbsInt(1))
	fmt.Println(utils.AbsInt(77))
	fmt.Println(utils.AbsInt(-77))

	// Output:
	// 0
	// 1
	// 1
	// 77
	// 77
}
