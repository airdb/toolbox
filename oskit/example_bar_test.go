package oskit_test

import (
	"time"

	"github.com/airdb/toolbox/oskit"
)

func ExampleBar_NewOption() {
	var bar oskit.Bar

	bar.NewOption(0, 100)

	for i := 0; i <= 100; i++ {
		time.Sleep(time.Millisecond)
		bar.Play(int64(i))
	}

	bar.Finish()
}
