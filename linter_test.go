package nilnesserr_test

import (
	"testing"

	"github.com/alingse/nilnesserr"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc     string
		settings nilnesserr.LinterSetting
	}{
		{
			desc:     "nilnesserr",
			settings: nilnesserr.LinterSetting{},
		},
		{
			desc:     "cgo",
			settings: nilnesserr.LinterSetting{},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			a, err := nilnesserr.NewAnalyzer(test.settings)
			if err != nil {
				t.Fatal(err)
			}

			analysistest.Run(t, analysistest.TestData(), a, test.desc)
		})
	}
}
