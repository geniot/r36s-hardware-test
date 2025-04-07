package main

import (
	"embed"
	"github.com/veandco/go-sdl2/sdl"
)

type ButtonCode = uint8
type ResourceKey = int

const (
	APP_NAME    = "R36S Hardware Test"
	APP_VERSION = "0.1"
)

var (
	//go:embed media/*
	mediaList embed.FS
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

	BACKGROUND_COLOR = COLOR_BLACK
)

const (
	RESOURCE_BGR_KEY           = ResourceKey(iota)
	RESOURCE_CIRCLE_YELLOW_KEY = ResourceKey(iota)
	RESOURCE_CROSS_YELLOW_KEY  = ResourceKey(iota)
	RESOURCE_CIRCLE_RED_KEY    = ResourceKey(iota)
)

const (
	BUTTON_CODE_FN             = ButtonCode(12)
	BUTTON_CODE_START          = ButtonCode(13)
	BUTTON_CODE_LEFT_JOYSTICK  = ButtonCode(14)
	BUTTON_CODE_RIGHT_JOYSTICK = ButtonCode(15)
	BUTTON_CODE_SELECT         = ButtonCode(16)
)

const (
	SCREEN_LEFT_UP_X    = int32(235)
	SCREEN_LEFT_UP_Y    = int32(130)
	SCREEN_RIGHT_DOWN_X = int32(400)
	SCREEN_RIGHT_DOWN_Y = int32(253)
	SCREEN_WIDTH        = SCREEN_RIGHT_DOWN_X - SCREEN_LEFT_UP_X
	SCREEN_HEIGHT       = SCREEN_RIGHT_DOWN_Y - SCREEN_LEFT_UP_Y
)

var ()
