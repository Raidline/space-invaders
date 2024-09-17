package main

import (
	"raidline/space-invaders/pkg/engine"
	"raidline/space-invaders/pkg/window"
)

func main() {
	c := window.Make()
	eng := engine.Make(c)
	eng.Run()
}
