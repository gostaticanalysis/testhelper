package main

import (
	"github.com/gostaticanalysis/testhelper"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(testhelper.Analyzer) }
