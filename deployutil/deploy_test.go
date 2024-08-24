package deployutil_test

import (
	"testing"

	"github.com/airdb/toolbox/deployutil"
	. "github.com/smartystreets/goconvey/convey"
)

func TestToEnv(t *testing.T) {
	Convey("Given array with few string value", t, func() {
		So(deployutil.ToEnv("blue"), ShouldEqual, deployutil.DeployPolicyBlue)
	})
}
