package main

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/singlechecker"
)

var DumbLinter = &analysis.Analyzer{
	Name: "dumb_linter",
	Doc:  "dumb linter always reporting error on first call expression",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		wasReport := false
		ast.Inspect(file, func(n ast.Node) bool {
			if wasReport {
				return true
			}
			ce, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			pass.Reportf(ce.Pos(), "random error")
			wasReport = true
			return true

		})
	}

	return nil, nil
}

func main() {
	singlechecker.Main(DumbLinter)
}

type analyzerPlugin struct{}

// This must be implemented
func (*analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		DumbLinter,
	}
}

// This must be defined and named 'AnalyzerPlugin'
var AnalyzerPlugin analyzerPlugin
