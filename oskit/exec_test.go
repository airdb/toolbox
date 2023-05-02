package oskit_test

import (
	"testing"

	"github.com/airdb/toolbox/oskit"
)

func TestExec(t *testing.T) {
	bin := "echo"
	args := []string{"test"}

	oskit.Exec(bin, args)
}

func TestExecCommand(t *testing.T) {
	bin := "echo"
	args := []string{"test"}

	out, err := oskit.ExecCommand(bin, args)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(out)
}
