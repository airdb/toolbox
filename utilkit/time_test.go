package utilkit_test

import (
	"testing"
	"time"

	"github.com/airdb/toolbox/utilkit"
)

func TestGetNowLong(t *testing.T) {
	utilkit.GetNowLong()
}

func TestGetYearMonthDay(t *testing.T) {
	utilkit.GetYearMonthDay()
}

func TestGetTimeRFC(t *testing.T) {
	now := time.Now().Unix()
	utilkit.GetTimeRFC(now)
}
