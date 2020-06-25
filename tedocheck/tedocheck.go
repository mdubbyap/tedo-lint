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
	Run:  Run,
}

func Run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		if !strings.HasSuffix(pass.Fset.File(file.Pos()).Name(), "_test.go") {
			ast.Inspect(file, func(n ast.Node) bool {
				if s, ok := n.(*ast.SelectorExpr); ok {
					if x, ok := s.X.(*ast.Ident); ok && x.Name == "fmt" && (s.Sel.Name == "Println" || s.Sel.Name == "Printf" || s.Sel.Name == "Print") {
						pass.Reportf(s.Pos(), "output to console in non-test file found %q", render(pass.Fset, s))
					}
				}
				return true
			})
		}
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
