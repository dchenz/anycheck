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
	issues := []Issue{}
	for _, file := range pass.Files {
		ast.Inspect(file, func(node ast.Node) bool {
			if newIssue := checkForAny(node, cfg); newIssue != nil {
				issues = append(issues, newIssue)
			}
			if newIssue := checkForInterface(node, cfg); newIssue != nil {
				issues = append(issues, newIssue)
			}
			return true
		})
	}
	for _, issue := range issues {
		pass.Report(analysis.Diagnostic{
			Pos:      issue.Pos(),
			Category: "anycheck",
			Message:  issue.Message(),
		})
	}
	return nil, nil
}

func checkForAny(node ast.Node, cfg *config) Issue {
	switch n := node.(type) {
	case *ast.Ident:
		if n.Name == "any" && !cfg.allowAny {
			return anyNotAllowed{pos: node.Pos()}
		}
	}
	return nil
}

func checkForInterface(node ast.Node, cfg *config) Issue {
	switch n := node.(type) {
	case *ast.InterfaceType:
		// This is an empty interface with no methods.
		if len(n.Methods.List) == 0 && !cfg.allowInterface {
			return interfaceNotAllowed{pos: node.Pos()}
		}
	}
	return nil
}
