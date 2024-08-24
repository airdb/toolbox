package osutil_test

import (
	"testing"

	"github.com/airdb/toolbox/osutil"
)

func TestExec(t *testing.T) {
	bin := "echo"
	args := []string{"test"}

	osutil.Exec(bin, args)
}

func TestExecCommand(t *testing.T) {
	bin := "echo"
	args := []string{"test"}

	out, err := osutil.ExecCommand(bin, args)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(out)
}
