package osutil_test

import (
	"testing"

	"github.com/airdb/toolbox/osutil"
	. "github.com/smartystreets/goconvey/convey"
)

func TestToEnv(t *testing.T) {
	Convey("Given array with few string value", t, func() {
		So(osutil.ToEnv("blue"), ShouldEqual, osutil.DeployPolicyBlue)
	})
}
