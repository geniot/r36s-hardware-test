package main

import "github.com/veandco/go-sdl2/sdl"

const (
	APP_NAME    = "R36S Hardware Test"
	APP_VERSION = "0.1"
)

var (
	COLOR_RED    = sdl.Color{R: 192, G: 64, B: 64, A: 255}
	COLOR_GREEN  = sdl.Color{R: 64, G: 192, B: 64, A: 255}
	COLOR_GRAY   = sdl.Color{R: 192, G: 192, B: 192, A: 255}
	COLOR_WHITE  = sdl.Color{R: 255, G: 255, B: 255, A: 255}
	COLOR_PURPLE = sdl.Color{R: 255, G: 0, B: 255, A: 255}
	COLOR_YELLOW = sdl.Color{R: 255, G: 255, B: 0, A: 255}
	COLOR_BLUE   = sdl.Color{R: 0, G: 255, B: 255, A: 255}
	COLOR_BLACK  = sdl.Color{R: 0, G: 0, B: 0, A: 255}
)
