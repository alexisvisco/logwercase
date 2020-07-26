package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/alexisvisco/logwercase/internal/linter"
)

func main() {
	singlechecker.Main(linter.Analyzer)
}
