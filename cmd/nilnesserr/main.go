package main

import (
	"log"

	"github.com/alingse/nilnesserr"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	setting := nilnesserr.LinterSetting{}

	analyzer, err := nilnesserr.NewAnalyzer(setting)
	if err != nil {
		log.Fatal(err)
	}

	singlechecker.Main(analyzer)
}
