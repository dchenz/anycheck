package analyzer

import (
	"fmt"
	"go/ast"

	"github.com/dchenz/anycheck/anycheck"
	"golang.org/x/tools/go/analysis"
)

const (
	vAny       = "any"
	vInterface = "interface"
)

type LinterSettings struct {
	Prefer string
}

type analyzer struct {
	allowAny       bool
	allowInterface bool
}

func NewAnalyzer(settings LinterSettings) (*analysis.Analyzer, error) {
	a := analyzer{}

	if settings.Prefer == "" {
		settings.Prefer = vAny
	}

	switch settings.Prefer {
	case vAny:
		a.allowAny = true
	case vInterface:
		a.allowInterface = true
	default:
		return nil, fmt.Errorf("expected '%s' or '%s' as preference", vAny, vInterface)
	}

	return &analysis.Analyzer{
		Name: "anycheck",
		Doc:  "checks the usage of 'interface{}' and 'any' in go code",
		Run:  a.run,
	}, nil
}

func (a *analyzer) run(pass *analysis.Pass) (interface{}, error) {
	linter, err := anycheck.NewLinter(
		anycheck.SetAllowAny(a.allowAny),
		anycheck.SetAllowInterface(a.allowInterface),
	)
	if err != nil {
		return nil, err
	}
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
