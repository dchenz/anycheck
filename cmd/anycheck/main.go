package main

import (
	anycheck "github.com/dchenz/anycheck/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(anycheck.NewAnalyzer())
}
