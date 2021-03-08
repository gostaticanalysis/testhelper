package testhelper

import (
	"go/types"
	"strings"

	"github.com/gostaticanalysis/analysisutil"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/ssa"
)

const doc = "testhelper finds a package function which is not a test function and receives a value of *testing.T as a parameter but it does not call (*testing.T).Helper"

// Analyzer finds a package function which is not a test function and receives a value of *testing.T as a parameter but it does not call (*testing.T).Helper.
var Analyzer = &analysis.Analyzer{
	Name: "testhelper",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		buildssa.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	pkg := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA).Pkg
	if strings.HasSuffix(pkg.Pkg.Path(), ".test") {
		return nil, nil
	}

	testingT := analysisutil.ObjectOf(pass, "testing", "T")
	if testingT == nil {
		return nil, nil
	}

	testingTPtr := types.NewPointer(testingT.Type())
	helperObj, _, _ := types.LookupFieldOrMethod(testingT.Type(), true, testingT.Pkg(), "Helper")
	helper, _ := helperObj.(*types.Func)
	if helper == nil {
		return nil, nil
	}

	for _, mem := range pkg.Members {
		fun, _ := mem.(*ssa.Function)
		if fun == nil || len(fun.Blocks) == 0 {
			continue
		}

		file := pass.Fset.File(fun.Pos())
		if file == nil {
			continue
		}

		// skip a test function
		if strings.HasSuffix(file.Name(), "_test.go") &&
			strings.HasPrefix(fun.Name(), "Test") &&
			len(fun.Params) == 1 &&
			types.Identical(testingTPtr, fun.Params[0].Type()) {
			continue
		}

		if hasTestingT(fun.Params, testingTPtr) && !isCalled(fun, helper) {
			pass.Reportf(fun.Pos(), "%s is a test helper but it does not call t.Helper", fun.Name())
		}
	}

	return nil, nil
}

func isCalled(fun *ssa.Function, helper *types.Func) bool {
	for _, b := range fun.Blocks {
		for _, instr := range b.Instrs {
			if analysisutil.Called(instr, nil, helper) {
				return true
			}
		}
	}
	return false
}

func hasTestingT(params []*ssa.Parameter, testingTPtr types.Type) bool {
	for _, p := range params {
		if types.Identical(testingTPtr, p.Type()) {
			return true
		}
	}
	return false
}
