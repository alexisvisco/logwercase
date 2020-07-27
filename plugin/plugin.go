package main

import (
	"golang.org/x/tools/go/analysis"

	"github.com/alexisvisco/logwercase/internal/linter"
)

// cf: https://golangci-lint.run/contributing/new-linters/#how-to-add-a-private-linter-to-golangci-lint

type analyzerPlugin struct{}

// This must be implemented
func (*analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		linter.LogMessageAnalyzer,
		linter.WithFieldAnalyzer,
	}
}

// This must be defined and named 'AnalyzerPlugin'
var AnalyzerPlugin analyzerPlugin
