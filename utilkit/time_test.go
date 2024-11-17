package utilkit_test

import (
	"strings"
	"testing"
	"time"

	"github.com/airdb/toolbox/utilkit"
)

func TestGetNowLong(t *testing.T) {
	ts := utilkit.GetNowLong()
	if ts == "" {
		t.Errorf("Expected a valid timestamp, got %s", ts)
	}
}

func TestGetYearMonthDay(t *testing.T) {
	today := utilkit.GetYearMonthDay()
	if !strings.Contains(today, "-") {
		t.Errorf("Expected a valid date string, got %s", today)
	}

}

func TestGetTimeRFC(t *testing.T) {
	now := time.Now().Unix()
	rfcTime := utilkit.GetTimeRFC(now)
	if _, err := time.Parse(time.RFC3339, rfcTime); err != nil {
		t.Errorf("Expected valid RFC3339 time format, got %s", rfcTime)
	}
}
