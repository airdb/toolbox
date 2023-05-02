package oskit_test

import (
	"testing"
	"time"

	"github.com/airdb/toolbox/oskit"
)

func TestProcessBar(t *testing.T) {
	var bar oskit.Bar

	bar.NewOption(0, 100)

	for i := 0; i <= 100; i++ {
		time.Sleep(10 * time.Millisecond)
		bar.Play(int64(i))
	}

	bar.Finish()
}
