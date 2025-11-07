package utils /* Utils */

import (
	"reflect"
	"runtime"
	"strings"
)

// TODO docs
func GetFunctionName(temp interface{}) string {
	strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name()), ".")
	return strs[len(strs)-1]
}

// #
// Ternary operator. The same as "(map[bool]T {true: vTrue, false: vFalse})[cond]".
// #
func TerOp[T any](cond bool, vTrue, vFalse T) T {
	if cond {
		return vTrue
	} else {
		return vFalse
	}
}
