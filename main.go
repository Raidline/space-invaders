package main

import (
	"raidline/space-invaders/pkg/window"
	"time"
)

func main() {
	c := window.Make(50, 120)
	for {
		c.Draw()
		time.Sleep(time.Millisecond * 500)
	}
}
