package main

import (
	"log"

	anycheck "github.com/dchenz/anycheck/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	a, err := anycheck.NewAnalyzer(anycheck.LinterSettings{})
	if err != nil {
		log.Fatalf("anycheck error: %s", err.Error())
	}
	singlechecker.Main(a)
}
