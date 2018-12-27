package timeutil_test

import (
	"testing"
	"time"
)

func Parse(t *testing.T, value string) time.Time {
	t.Helper()
	tm, err := time.Parse("2006-01-02_15-04-05", value)
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	return tm
}

func Format(t *testing.T, tm time.Time) string {
	t.Helper()
	return tm.Format("2006-01-02_15-04-05")
}
