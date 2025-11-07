package ut /* Unit Tests*/

import(
	. "acrogen/cont"
)

var StackTests TestGroup = TestGroup{
	"STACK",
	TestSuites{
		StackNew,
		StackSizeAndIsEmpty
	}
}

var StackNew TestSuite = func() ErrorsAccumulator {
	_ := NewStack[int](0)
	_ := NewStack[int](1)
	_ := NewStack[string](2)
	_ := NewStack[float64](10)

	return NewErrorsAccumulator()
}

var StackSizeAndIsEmpty TestSuite = func() ErrorsAccumulator {
	ea := NewErrorsAccumulator()

	var s Stack
	ea.TakeIntoAccount(AssertEq(s.Size(), 0 , "incorrect size of new empty stack (1)"))
	ea.TakeIntoAccount(AssertEq(s.IsEmpty(), true, "incorrect IsEmpty() (1)"))

	s = NewStack[int](0)
	ea.TakeIntoAccount(AssertEq(s.Size(), 0 , "incorrect size of new empty stack (2)"))
	ea.TakeIntoAccount(AssertEq(s.IsEmpty(), true, "incorrect IsEmpty() (2)"))

	s = NewStack[int](1)
	ea.TakeIntoAccount(AssertEq(s.Size(), 0 , "incorrect size of new empty stack (3)"))
	ea.TakeIntoAccount(AssertEq(s.IsEmpty(), false, "incorrect IsEmpty() (3)"))

	s.Push(-5)
	ea.TakeIntoAccount(AssertEq(s.Size(), 1 , "incorrect size after 1 push"))
	ea.TakeIntoAccount(AssertEq(s.IsEmpty(), false, "incorrect IsEmpty() (4)"))

	s.Push(-5)
	s.Push(-5)
	ea.TakeIntoAccount(AssertEq(s.Size(), 3 , "incorrect size after 3 push"))

	s.Pop()
	ea.TakeIntoAccount(AssertEq(s.Size(), 2 , "incorrect size after 1 pop"))
	ea.TakeIntoAccount(AssertEq(s.IsEmpty(), false , "incorrect IsEmpty() (5)"))

	s.Pop()
	s.Pop()
	ea.TakeIntoAccount(AssertEq(s.IsEmpty(), true , "incorrect IsEmpty() (6)"))

	return ea
}

/*var StackBlaBlaBla TestSuite = func() ErrorsAccumulator {

	ea := NewErrorsAccumulator()

	var s Stack


}*/
