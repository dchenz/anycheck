package anycheck

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "anycheck",
		Doc:  "checks the usage of 'interface{}' and 'any' in go code",
		Run:  run,
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	cfg := &config{}
	for _, file := range pass.Files {
		ast.Inspect(file, func(node ast.Node) bool {
			return checkNode(node, pass, cfg)
		})
	}
	return nil, nil
}

func checkNode(node ast.Node, pass *analysis.Pass, cfg *config) bool {
	switch n := node.(type) {
	case *ast.InterfaceType:
		// This is an empty interface with no methods.
		if len(n.Methods.List) == 0 && !cfg.allowInterface {
			reportInvalidInterface(node, pass)
		}
	}
	return true
}

func reportInvalidInterface(node ast.Node, pass *analysis.Pass) {
	pass.Report(analysis.Diagnostic{
		Pos:      node.Pos(),
		End:      node.End(),
		Category: "anycheck",
		Message:  "interface detected",
	})
}
