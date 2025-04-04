package logger

import (
	"fmt"
	"log"
	"raidline/space-invaders/pkg/colors"
)

const errorPrefix = "space Invaders had a fatal error %s"

func colorize(color, message string) string {
	return fmt.Sprintf("%s%s%s", color, message, colors.Reset)
}

func Error(msg string, values ...interface{}) {
	errMsg := fmt.Sprintf(errorPrefix, fmt.Sprintf(msg, values))
	log.Println(colorize(colors.Red, errMsg))
}

func Warn(msg string) {
	log.Println(colorize(colors.Blue, msg))
}

func Debug(msg string, values ...interface{}) {
	log.Println(colorize(colors.Green, fmt.Sprintf(msg, values)))
}
