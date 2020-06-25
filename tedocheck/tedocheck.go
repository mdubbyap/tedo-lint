package tedocheck

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "tedolint",
	Doc:  "reports tedo things",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			if strings.HasSuffix(file.Name.Name, "_test.go") {
				return true
			}
			if s, ok := n.(*ast.SelectorExpr); ok {
				if x, ok := s.X.(*ast.Ident); ok && x.Name == "fmt" && (s.Sel.Name == "Println" || s.Sel.Name == "Printf" || s.Sel.Name == "Print") {
					pass.Reportf(s.Pos(), "output to console in non-test file found %q", render(pass.Fset, s))
				}
			}
			return true

			//if !ok {
			//	return true
			//}
			//
			//if be.Op != token.ADD {
			//	return true
			//}
			//
			//if _, ok := be.X.(*ast.BasicLit); !ok {
			//	return true
			//}
			//
			//if _, ok := be.Y.(*ast.BasicLit); !ok {
			//	return true
			//}
			//
			//isInteger := func(expr ast.Expr) bool {
			//	t := pass.TypesInfo.TypeOf(expr)
			//	if t == nil {
			//		return false
			//	}
			//
			//	bt, ok := t.Underlying().(*types.Basic)
			//	if !ok {
			//		return false
			//	}
			//
			//	if (bt.Info() & types.IsInteger) == 0 {
			//		return false
			//	}
			//
			//	return true
			//}
			//
			//// check that both left and right hand side are integers
			//if !isInteger(be.X) || !isInteger(be.Y) {
			//	return true
			//}
			//
			//pass.Reportf(be.Pos(), "integer addition found %q",
			//	render(pass.Fset, be))
			//return true
		})
	}

	return nil, nil
}

// render returns the pretty-print of the given node
func render(fset *token.FileSet, x interface{}) string {
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, fset, x); err != nil {
		panic(err)
	}
	return buf.String()
}
