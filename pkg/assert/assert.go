package assert

import (
	"fmt"
	"os"
	"raidline/space-invaders/pkg/logger"
)

func NonNil(obj interface{}) {
	Assert(obj != nil, "%s should not be nil", obj)
}

func Assert(condition bool, msg string, values ...interface{}) {
	if !condition {

		_, _ = fmt.Fprint(os.Stdout, "\x1b[H\x1b[2J\x1b[?25h\x1b[0m")

		var formattedMsg string

		if len(values) != 0 {
			formattedMsg = fmt.Sprintf(msg, values)
		} else {
			formattedMsg = msg
		}

		logger.Error(formattedMsg)

		panic(formattedMsg)
	}
}
