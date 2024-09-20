package engine

import (
	"raidline/space-invaders/game"
	"raidline/space-invaders/pkg/window"
	"time"
)

const MaxSamples = 100

type stats struct {
	start time.Time
	fps   float64
}

func (e *Engine) calculateFps() {
	if e.frames == MaxSamples {
		e.stats.fps = float64(e.frames) / time.Since(e.stats.start).Seconds()
		e.frames = 0
		e.stats.start = time.Now()
	}
}

type Engine struct {
	frames  int
	canvas  *window.Canvas
	game    *game.Game
	stats   stats
	running bool
}

func Make(w *window.Canvas, g *game.Game) *Engine {
	return &Engine{
		frames: 0,
		canvas: w,
		game:   g,
		stats: stats{
			start: time.Now(),
			fps:   0,
		},
		running: true,
	}
}

func (e *Engine) Run() {
	e.canvas.Flush()
	for e.running {
		e.frames++
		e.calculateFps()
		e.canvas.Draw(e.stats.fps, e.game.Board)
	}
}
