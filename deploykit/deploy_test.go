package deploykit_test

import (
	"testing"

	"github.com/airdb/toolbox/deploykit"
	. "github.com/smartystreets/goconvey/convey"
)

func TestToEnv(t *testing.T) {
	Convey("Given array with few string value", t, func() {
		So(deploykit.ToEnv("blue"), ShouldEqual, deploykit.DeployPolicyBlue)
	})
}
