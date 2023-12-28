package main

import (
	"flag"
	"log"

	anycheck "github.com/dchenz/anycheck/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

var settings anycheck.LinterSettings

func main() {
	flag.Parse()

	a, err := anycheck.NewAnalyzer(settings)
	if err != nil {
		log.Fatalf("anycheck error: %s", err.Error())
	}
	singlechecker.Main(a)
}

func init() {
	flag.StringVar(&settings.Prefer, "prefer", "any",
		"linter preference: either 'any' or 'interface'")
}
