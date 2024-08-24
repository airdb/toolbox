package oskit_test

import (
	"testing"

	"github.com/airdb/toolbox/osutil"
	. "github.com/smartystreets/goconvey/convey"
)

func TestOS(t *testing.T) {
	Convey("Given array with few string value", t, func() {
		Convey("Then reverse the array", func() {
			t.Log("win:", osutil.IsWin())
			t.Log("mac:", osutil.IsMac())
			t.Log("linux:", osutil.IsLinux())
		})
	})
}
