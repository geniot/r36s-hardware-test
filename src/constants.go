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

var (
	RESOURCE_BGR_KEY           = 0
	RESOURCE_CIRCLE_YELLOW_KEY = 1
	RESOURCE_CROSS_YELLOW_KEY  = 2
)

var (
	BUTTON_CODE_SELECT = uint8(16)
	BUTTON_CODE_START  = uint8(13)
)

var (
	SCREEN_LEFT_UP_X    = int32(235)
	SCREEN_LEFT_UP_Y    = int32(130)
	SCREEN_RIGHT_DOWN_X = int32(400)
	SCREEN_RIGHT_DOWN_Y = int32(253)
	SCREEN_WIDTH        = SCREEN_RIGHT_DOWN_X - SCREEN_LEFT_UP_X
	SCREEN_HEIGHT       = SCREEN_RIGHT_DOWN_Y - SCREEN_LEFT_UP_Y
)
