package cont

import (
	"fmt"
)

func ExampleNewIdPermutation() {

	p := NewIdPermutation(0)
	fmt.Println(p.elems)

	p = NewIdPermutation(1)
	fmt.Println(p.elems)

	p = NewIdPermutation(2)
	fmt.Println(p.elems)

	p = NewIdPermutation(8)
	fmt.Println(p.elems)

	// Output:
	// []
	// [0]
	// [0 1]
	// [0 1 2 3 4 5 6 7]
}

func ExamplePermutation_Len() {

	fmt.Println(NewIdPermutation(0).Len())
	fmt.Println(NewIdPermutation(1).Len())
	fmt.Println(NewIdPermutation(2).Len())
	fmt.Println(NewIdPermutation(8).Len())

	// Output:
	// 0
	// 1
	// 2
	// 8
}

func ExamplePermutation_Resize() {

	p := NewIdPermutation(3)
	p.Resize(5)
	fmt.Println(p.elems)

	// Output:
	// [0 1 2 3 4]
}

func ExamplePermutation_Shift() {

	p := NewIdPermutation(3)

	p.Shift(0)
	fmt.Println(p.elems)

	p.Shift(1)
	fmt.Println(p.elems)

	p.Shift(1)
	fmt.Println(p.elems)

	p.Shift(-2)
	fmt.Println(p.elems)

	p.Shift(1 * 2 * 3)
	fmt.Println(p.elems)

	// Output:
	// [0 1 2]
	// [0 2 1]
	// [1 0 2]
	// [0 1 2]
	// [0 1 2]
}

func ExamplePermutation_Next() {

	p := NewIdPermutation(3)
	p = p.Next().Next()
	fmt.Println(p.elems)

	// Output:
	// [1 0 2]
}

func ExamplePermutation_Prev() {

	p := NewIdPermutation(3)

	p = p.Prev()
	fmt.Println(p.elems)

	p.Shift(2)
	p = p.Prev()
	fmt.Println(p.elems)

	// Output:
	// [2 1 0]
	// [0 1 2]
}

func ExamplePermutation_Get() {

	p := NewIdPermutation(3)
	p.Shift(2)

	fmt.Println(p.Get(0))
	fmt.Println(p.Get(1))
	fmt.Println(p.Get(2))

	// Output:
	// 1
	// 0
	// 2
}

func ExamplePermutation_Clone() {

	p := NewIdPermutation(3)
	p.Shift(2)

	pc := p.Clone()
	fmt.Println(pc.elems)

	pc.Shift(-2)
	fmt.Println(pc.elems)

	fmt.Println(p.elems)

	// Output:
	// [1 0 2]
	// [0 1 2]
	// [1 0 2]
}

func ExamplePermutation_AsSlice() {

	p := NewIdPermutation(3)
	p.Shift(2)
	fmt.Println(p.AsSlice())

	// Output:
	// [1 0 2]
}

func ExampleIsPermutation() {

	slice := []int{}
	fmt.Println(IsPermutation(slice)) // false

	slice = []int{0}
	fmt.Println(IsPermutation(slice)) // true

	slice = []int{1}
	fmt.Println(IsPermutation(slice)) // false

	slice = []int{1, 2, 0, 5, 4, 3}
	fmt.Println(IsPermutation(slice)) // true

	slice = []int{1, 2, 0, 8, 4, 3}
	fmt.Println(IsPermutation(slice)) // false

	slice = []int{0, 0, 0}
	fmt.Println(IsPermutation(slice)) // false

	// Output:
	// false
	// true
	// false
	// true
	// false
	// false
}

func ExampleGetPermutatedSlice() {

	slice := []int{1, 2, 3, 4, 5}
	p := NewIdPermutation(5)
	slice, _ = GetPermutatedSlice(slice, p)
	fmt.Println(slice) // [1 2 3 4 5]

	slice = []int{1, 2, 3, 4, 5}
	p = NewIdPermutation(2)
	slice, _ = GetPermutatedSlice(slice, p)
	fmt.Println(slice) // [1 2 3 4 5]

	slice = []int{1, 2, 3, 4, 5}
	p = NewIdPermutation(2)
	p.Shift(1)
	slice, _ = GetPermutatedSlice(slice, p)
	fmt.Println(slice) // [2 1 3 4 5]

	slice = []int{1, 2, 3, 4, 5}
	p = NewIdPermutation(5)
	p.Shift(8)
	slice, _ = GetPermutatedSlice(slice, p)
	fmt.Println(slice) // [1 3 4 2 5]

	slice = []int{1, 2, 3, 4, 5}
	p = NewIdPermutation(5)
	p.Shift(2)
	slice, _ = GetPermutatedSlice(slice, p)
	fmt.Println(slice) // [1 2 4 3 5]

	sliceStr := []string{"A", "B"}
	p = NewIdPermutation(2)
	p.Shift(1)
	sliceStr, _ = GetPermutatedSlice(sliceStr, p)
	fmt.Println(sliceStr) // [B A]

	sliceStr = []string{"A"}
	p = NewIdPermutation(1)
	p.Shift(1)
	sliceStr, _ = GetPermutatedSlice(sliceStr, p)
	fmt.Println(sliceStr) // [A]

	// Output:
	// [1 2 3 4 5]
	// [1 2 3 4 5]
	// [2 1 3 4 5]
	// [1 3 4 2 5]
	// [1 2 4 3 5]
	// [B A]
	// [A]
}

func ExamplePermutationsGroupOrder() {

	fmt.Println(PermutationsGroupOrder(0))
	fmt.Println(PermutationsGroupOrder(1))
	fmt.Println(PermutationsGroupOrder(2))
	fmt.Println(PermutationsGroupOrder(5))

	// Output:
	// 1
	// 1
	// 2
	// 120
}
