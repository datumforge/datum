package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	atolinter "github.com/datumforge/datum/pkg/ato-linter"
)

func main() {
	singlechecker.Main(atolinter.ATOAnalyzer)
}
