package linter

import (
	"go/ast"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

type logMessageAnalyzer struct{}

func newLogMessageAnalyzer() *logMessageAnalyzer {
	return &logMessageAnalyzer{}
}

var LogMessageAnalyzer = &analysis.Analyzer{
	Name: "logmessage",
	Doc:  "reports when first character of a log is in uppercase (it should be in lower)",
	Run:  newLogMessageAnalyzer().run,
}

var loggerTriggerName = map[string]bool{
	"log":       true,
	"logger":    true,
	"logrus":    true,
	"WithField": true,
	"WithErr":   true,
}

var loggerFunctionNames = map[string]bool{
	"Trace":     true,
	"Debug":     true,
	"Print":     true,
	"Info":      true,
	"Warn":      true,
	"Warning":   true,
	"Error":     true,
	"Panic":     true,
	"Fatal":     true,
	"Tracef":    true,
	"Debugf":    true,
	"Printf":    true,
	"Infof":     true,
	"Warnf":     true,
	"Warningf":  true,
	"Errorf":    true,
	"Panicf":    true,
	"Fatalf":    true,
	"Traceln":   true,
	"Debugln":   true,
	"Println":   true,
	"Infoln":    true,
	"Warnln":    true,
	"Warningln": true,
	"Errorln":   true,
	"Panicln":   true,
	"Fatalln":   true,
}

func (l logMessageAnalyzer) run(pass *analysis.Pass) (interface{}, error) {
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

			if _, ok := loggerFunctionNames[function.Sel.Name]; ok &&
				l.isLoggerCall(callExpr) &&
				l.isCapitalized(callExpr.Args) {
				pass.ReportRangef(callExpr.Args[0], "Log message should be lower cased like that: %q",
					l.toLowerCase(callExpr.Args))
			}
			return true
		})
	}

	return nil, nil
}

func (l logMessageAnalyzer) isLoggerCall(expr *ast.CallExpr) bool {
	s, ok := expr.Fun.(*ast.SelectorExpr)
	if ok {
		if _, ok := loggerTriggerName[s.Sel.Name]; ok {
			return true
		}

		if s.X != nil {
			id, ok := s.X.(*ast.Ident)
			if ok {
				if _, ok := loggerTriggerName[id.Name]; ok {
					return true
				}
			}
			if e, ok := s.X.(*ast.CallExpr); ok {
				return l.isLoggerCall(e)
			}
			return false
		}
	}

	return false
}

func (l logMessageAnalyzer) isCapitalized(args []ast.Expr) bool {
	if len(args) == 0 {
		return false
	}

	lit, ok := args[0].(*ast.BasicLit)
	if !ok {
		return false
	}

	literalString := strings.Trim(lit.Value, "\"")
	if len(literalString) > 0 {
		return unicode.IsUpper(rune(literalString[0]))
	}

	return false
}

func (l logMessageAnalyzer) toLowerCase(args []ast.Expr) string {
	if len(args) == 0 {
		return ""
	}

	lit, ok := args[0].(*ast.BasicLit)
	if !ok {
		return ""
	}

	literalString := strings.Trim(lit.Value, "\"")

	if len(literalString) > 0 {
		firstChar := unicode.ToLower(rune(literalString[0]))
		return string([]rune{firstChar}) + literalString[1:]
	}

	return ""
}
