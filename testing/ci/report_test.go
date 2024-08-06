package ci

import (
	"os"
	"testing"
)

var (
	isCi = os.Getenv("CI") != ""
)

func TestMonthlyReport(t *testing.T) {
	if !isCi {
		t.Skip("not in CI system")
	}
	// test code
	t.Log("running in CI")
	// TODO: Rest of long running test ...
}
