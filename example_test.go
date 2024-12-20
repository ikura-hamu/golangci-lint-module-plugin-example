package golangci_lint_module_plugin_example_test

import (
	"testing"

	"github.com/gostaticanalysis/testutil"
	"github.com/ikura-hamu/golangci_lint_module_plugin_example"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, golangci_lint_module_plugin_example.Analyzer, "a")
}
