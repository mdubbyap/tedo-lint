package main

import (
	"github.com/mdubbyap/tedo-lint/tedocheck"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(tedocheck.Analyzer)
}
