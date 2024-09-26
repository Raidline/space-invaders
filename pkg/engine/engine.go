package engine

import (
	"raidline/space-invaders/game"
	"raidline/space-invaders/pkg/assert"
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
	frames     int
	canvas     *window.Canvas
	game       *game.Game
	stats      stats
	running    bool
	FramesChan chan bool
	statesChan <-chan int
}

func Make(w *window.Canvas) *Engine {
	framesChan := make(chan bool)
	e := &Engine{
		frames: 0,
		canvas: w,
		stats: stats{
			start: time.Now(),
			fps:   0,
		},
		running:    true,
		FramesChan: framesChan,
	}

	return e
}

func (e *Engine) SetGame(g *game.Game) {
	e.game = g
	e.statesChan = g.GameStatus
}

func (e *Engine) Run() {
	assert.NonNil(e.game)
	e.canvas.Flush()

	go func() {
		for {
			st := <-e.statesChan

			if st == game.END {
				e.running = false

				break
			}
		}
	}()

	defer close(e.FramesChan)
	for e.running {
		e.frames++
		e.FramesChan <- true
		e.calculateFps()
		e.canvas.Draw(e.stats.fps, e.game.Board)
	}
}
