package main

import (
	"golang.org/x/tools/go/analysis/multichecker"

	"github.com/alexisvisco/logwercase/internal/linter"
)

func main() {
	multichecker.Main(
		linter.AnalyzerMessageCase,
		linter.AnalyzerWithFieldCase)
}
