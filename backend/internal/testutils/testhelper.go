package testutils

import (
	"testing"
)

var (
	greenColor  = "\033[32m"
	redColor    = "\033[31m"
	resetColor  = "\033[0m"
	yellowColor = "\033[33m"
)

type TestStats struct {
	Total   int
	Passed  int
	Failed  int
	Package string
}

func NewTestStats(packageName string) *TestStats {
	return &TestStats{Package: packageName}
}

func (ts *TestStats) LogSuccess(t *testing.T, message string) {
	t.Logf("Testing: %s", message)
	t.Logf("%s✓ PASS: %s%s", greenColor, message, resetColor)
	ts.Passed++
	ts.Total++
}

func (ts *TestStats) LogError(t *testing.T, message string) {
	t.Logf("Testing: %s", message)
	t.Errorf("%s✗ FAIL: %s%s", redColor, message, resetColor)
	ts.Failed++
	ts.Total++
}

func (ts *TestStats) LogInfo(t *testing.T, message string) {
	t.Logf("%s• INFO: %s%s", yellowColor, message, resetColor)
}

func (ts *TestStats) PrintSummary(t *testing.T) {
	t.Logf("\n%s=== Test Summary for %s ===%s", yellowColor, ts.Package, resetColor)
	t.Logf("%s=== Total Tests: %d ===%s", yellowColor, ts.Total, resetColor)
	t.Logf("%s=== Passed: %d ===%s", greenColor, ts.Passed, resetColor)
	if ts.Failed > 0 {
		t.Logf("%s=== Failed: %d ===%s", redColor, ts.Failed, resetColor)
	} else {
		t.Logf("=== Failed: %d ===", ts.Failed)
	}
	t.Logf("%s=== End Summary ===%s", yellowColor, resetColor)
}
