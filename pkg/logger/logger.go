package logger

import (
	"fmt"
	"log"
	"raidline/space-invaders/pkg/colors"
)

func colorize(color, message string) string {
	return fmt.Sprintf("%s%s%s", color, message, colors.Reset)
}

func Error(msg string, values ...interface{}) {
	log.Fatalf(colorize(colors.Red, "space Invaders had a fatal error %s"), fmt.Sprintf(msg, values))
}

func Debug(msg string, values ...interface{}) {
	log.Print(colorize(colors.Green, fmt.Sprintf(msg, values)))
}
