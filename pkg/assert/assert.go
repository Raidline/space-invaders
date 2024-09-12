package assert

import "raidline/space-invaders/pkg/logger"

func Assert(condition bool, msg string, values ...interface{}) {
	if !condition {
		logger.Error(msg)
		panic(msg)
	}
}
