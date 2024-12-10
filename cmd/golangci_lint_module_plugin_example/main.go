package main

import (
	"github.com/ikura-hamu/golangci_lint_module_plugin_example"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(golangci_lint_module_plugin_example.Analyzer) }
