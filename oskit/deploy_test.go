package oskit_test

import (
	"testing"

	"github.com/airdb/toolbox/oskit"
	. "github.com/smartystreets/goconvey/convey"
)

func TestToEnv(t *testing.T) {
	Convey("Given array with few string value", t, func() {
		So(oskit.ToEnv("blue"), ShouldEqual, oskit.DeployPolicyBlue)
	})
}
