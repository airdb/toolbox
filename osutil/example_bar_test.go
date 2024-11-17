package osutil_test

import (
	"time"

	"github.com/airdb/toolbox/osutil"
)

func ExampleBar_NewOption() {
	var bar osutil.Bar

	bar.NewOption(0, 100)

	for i := 0; i <= 100; i++ {
		time.Sleep(time.Millisecond)
		bar.Play(int64(i))
	}

	bar.Finish()
}
