package main

import (
	"fmt"
	"go/ast"
	"os"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/analysis/singlechecker"
	"golang.org/x/tools/go/ast/inspector"
)

var (
	forbidden = []string{
		"syscall",
		// TODO: More packages to forbid
	}

	// Analyzer is the main entry point for the analyzer
	Analyzer = &analysis.Analyzer{
		Name:     "forbidden",
		Doc:      "reports the use of forbidden packages",
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Run:      run,
	}
)

func isForbidden(path string) bool {
	for _, f := range forbidden {
		if strings.HasPrefix(path, f) {
			return true
		}
	}
	return false
}

func run(pass *analysis.Pass) (any, error) {
	filter := []ast.Node{
		(*ast.ImportSpec)(nil),
	}

	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	inspect.Preorder(filter, func(node ast.Node) {
		imp, ok := node.(*ast.ImportSpec)
		if !ok {
			pos := pass.Fset.Position(node.Pos())
			fname, lnum := pos.Filename, pos.Line
			const warnFmt = "%s:%d: warning: not an import\n"
			fmt.Fprintf(os.Stderr, warnFmt, fname, lnum)
			return
		}

		// `"syscall"` -> `syscall`
		path := strings.Trim(imp.Path.Value, `"`)
		if isForbidden(path) {
			return
		}
		pass.Reportf(imp.Pos(), "importing forbidden package %q", path)

	})

	return nil, nil

}

func main() {
	singlechecker.Main(Analyzer)
}
