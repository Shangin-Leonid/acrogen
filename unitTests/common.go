package ut /* Unit Tests*/

// TODO Color output

import (
	"acrogen/utils"
)

// TODO docs
type TestSuite func() ErrorsAccumulator
type TestSuites []TestSuite

// TODO docs
type TestGroup struct {
	header     string
	testSuites TestSuites
}
type TestGroups []TestGroup

// TODO docs
type ModuleTests struct {
	header     string
	testGroups TestGroups
}
type ModulesTests []ModuleTest

// TODO docs
func RunModulesTests(mst ModulesTest) {
	for i := range mst {
		RunModuleTests(mst[i])
		NeutralColor.Printf("\n")
	}
}

// TODO docs
func RunModuleTests(mt ModuleTest) {
	NeutralBoldColor.Printf("\n===== %s =====\n\n", mt.header)
	for i := range mt.testGroups {
		RunTestGroup(mt.testGroups[i])
		NeutralColor.Printf("\n")
	}
}

// TODO docs
func RunTestGroup(tm TestGroup) {
	var ea ErrorsAccumulator
	NeutralBoldColor.Printf("\n%s:\n", tm.header)
	for i := range tm.testSuites {
		ea = testSuites[i]()
		NeutralColor.Printf("%s: ", utils.GetFunctionName(testSuites[i]))
		if ea.IsNoError() {
			PassedColor.Printf("%s\n", PassedMes)
		} else {
			FailedColor.Printf("%s ", FailedMes)
			for i := range ea.errs {
				FailReasonColor.Printf("(%s)\n", ea.errs[i].Error())
			}
		}
	}
}
