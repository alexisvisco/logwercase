package linter

import (
	"go/ast"
	"strings"

	"github.com/iancoleman/strcase"
	"golang.org/x/tools/go/analysis"
)

type logWithFieldAnalyzer struct{}

func newWithFieldAnalyzer() *logWithFieldAnalyzer {
	return &logWithFieldAnalyzer{}
}

var AnalyzerWithFieldCase = &analysis.Analyzer{
	Name: "withfield",
	Doc:  "reports when first character of a log is in uppercase (it should be in lower)",
	Run:  newWithFieldAnalyzer().run,
}

func (l logWithFieldAnalyzer) run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			callExpr, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			function, ok := callExpr.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}

			if function.Sel.Name == "WithField" {

				if len(callExpr.Args) == 0 {
					return true
				}

				arg, ok := callExpr.Args[0].(*ast.BasicLit)
				if !ok || !strings.HasPrefix(arg.Value, "\"") {
					return true
				}

				if kc, isNotKebabCase := isNotKebabCase(arg.Value); isNotKebabCase {
					pass.ReportRangef(arg, "WithField key should be in lower kebab case like that: %q", kc)
				}
			}

			return true
		})
	}

	return nil, nil
}

func isNotKebabCase(key string) (kc string, isKebab bool) {
	trimmedKey := strings.Trim(key, "\"")

	kebabCased := strings.Trim(strcase.ToKebab(trimmedKey), "-")

	return kebabCased, trimmedKey != kebabCased
}
