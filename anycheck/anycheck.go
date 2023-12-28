package anycheck

import (
	"go/ast"
)

type Linter struct {
	cfg config
}

func NewLinter() *Linter {
	return &Linter{}
}

func (l *Linter) Run(nodes []ast.Node) []Issue {
	issues := []Issue{}
	for _, node := range nodes {
		if newIssue := l.checkForAny(node); newIssue != nil {
			issues = append(issues, newIssue)
		}
		if newIssue := l.checkForInterface(node); newIssue != nil {
			issues = append(issues, newIssue)
		}
	}
	return issues
}

func (l *Linter) checkForAny(node ast.Node) Issue {
	switch n := node.(type) {
	case *ast.Ident:
		if n.Name == "any" && !l.cfg.allowAny {
			return anyNotAllowed{pos: node.Pos()}
		}
	}
	return nil
}

func (l *Linter) checkForInterface(node ast.Node) Issue {
	switch n := node.(type) {
	case *ast.InterfaceType:
		// This is an empty interface with no methods.
		if len(n.Methods.List) == 0 && !l.cfg.allowInterface {
			return interfaceNotAllowed{pos: node.Pos()}
		}
	}
	return nil
}
