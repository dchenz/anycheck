package analyzer

import (
	"path/filepath"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestPreferInterface(t *testing.T) {
	a, err := NewAnalyzer(LinterSettings{
		Prefer: vInterface,
	})
	if err != nil {
		t.Fatal(err)
	}
	testdata := filepath.Join(analysistest.TestData(), "prefer_interface")
	analysistest.Run(t, testdata, a)
}
