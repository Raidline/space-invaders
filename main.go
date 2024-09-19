package main

import (
	"raidline/space-invaders/game"
	"raidline/space-invaders/pkg/engine"
	"raidline/space-invaders/pkg/logger"
	"raidline/space-invaders/pkg/window"
	"syscall"
	"unsafe"
)

const MaxRows = 128
const MaxCols = 128

type winsize struct {
	Row uint16
	Col uint16
}

func main() {
	ws := getWinsize()

	c := window.Make(ws.Row, ws.Col)
	g := game.Make(ws.Row, ws.Col)
	eng := engine.Make(c, g)
	eng.Run()
}

func getWinsize() *winsize {

	ws := &winsize{}
	retCode, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))

	if int(retCode) == -1 {
		logger.Error("Could not get terminal information")
		panic(errno)
	}

	if ws.Row <= 0 {
		ws.Row = MaxRows
	}

	if ws.Col <= 0 {
		ws.Col = MaxCols
	}

	return ws
}
