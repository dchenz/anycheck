package analyzer

import (
	"go/ast"

	"github.com/dchenz/anycheck/anycheck"
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
	linter := anycheck.NewLinter()
	for _, file := range pass.Files {
		issues := linter.Run(collectNodes(file))
		for _, issue := range issues {
			pass.Report(analysis.Diagnostic{
				Pos:      issue.Pos(),
				Category: "anycheck",
				Message:  issue.Message(),
			})
		}
	}
	return nil, nil
}

func collectNodes(file *ast.File) []ast.Node {
	nodes := []ast.Node{}
	ast.Inspect(file, func(n ast.Node) bool {
		nodes = append(nodes, n)
		return true
	})
	return nodes
}
