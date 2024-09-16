package assert

import (
	"fmt"
	"raidline/space-invaders/pkg/logger"
)

func Assert(condition bool, msg string, values ...interface{}) {
	if !condition {

		var formattedMsg string

		if len(values) != 0 {
			formattedMsg = fmt.Sprintf(msg, values)
		} else {
			formattedMsg = msg
		}

		logger.Error(formattedMsg)
		panic(msg)
	}
}
