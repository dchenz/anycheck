package main

import (
	"github.com/dchenz/anycheck/anycheck"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(anycheck.NewAnalyzer())
}
