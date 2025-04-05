package main

import (
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/tevino/abool/v2"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

type Application struct {
	settings         *Settings
	sdlWindow        *sdl.Window
	sdlRenderer      *sdl.Renderer
	pressedKeysCodes mapset.Set[sdl.Keycode]
	lastPressedKey   sdl.Keycode
	isRunning        *abool.AtomicBool
	backgroundColor  sdl.Color
}

func NewApplication() *Application {
	return &Application{
		pressedKeysCodes: mapset.NewSet[sdl.Keycode](),
		isRunning:        abool.New(),
		backgroundColor:  COLOR_GRAY}
}

func (app *Application) Start(args []string) {
	sdl.Init(sdl.INIT_EVERYTHING)
	app.settings = NewSettings()
	app.sdlWindow, _ = sdl.CreateWindow(
		APP_NAME+" "+APP_VERSION,
		int32(app.settings.WindowPosX), int32(app.settings.WindowPosY),
		int32(app.settings.WindowWidth), int32(app.settings.WindowHeight),
		uint32(app.settings.WindowState))
	app.sdlRenderer, _ = sdl.CreateRenderer(app.sdlWindow, -1, sdl.RENDERER_PRESENTVSYNC|sdl.RENDERER_ACCELERATED)
	app.isRunning.Set()
	for app.isRunning.IsSet() {
		app.UpdateEvents()
		app.UpdatePhysics()
		app.UpdateView()
	}
}

func (app *Application) Stop() {
	app.isRunning.UnSet()
	app.settings.Save(app.sdlWindow)
}

func (app *Application) UpdateEvents() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {

		case *sdl.KeyboardEvent:
			if t.Repeat > 0 {
				break
			}
			app.lastPressedKey = t.Keysym.Sym
			if t.State == sdl.PRESSED {
				app.pressedKeysCodes.Add(t.Keysym.Sym)
			} else { // if t.State == sdl.RELEASED {
				app.pressedKeysCodes.Remove(t.Keysym.Sym)
			}
			break

		case *sdl.WindowEvent:
			if t.Event == sdl.WINDOWEVENT_CLOSE {
				app.settings.SaveWindowState(app.sdlWindow)
			}
			break

		case *sdl.QuitEvent:
			app.Stop()
			break
		}
	}
}

func (app *Application) UpdatePhysics() {
	if app.pressedKeysCodes.Contains(sdl.K_q) {
		app.Stop()
	}
}

func (app *Application) UpdateView() {
	if err := app.sdlRenderer.SetDrawColorArray(app.backgroundColor.R, app.backgroundColor.G, app.backgroundColor.B, app.backgroundColor.A); err != nil {
		println(err.Error())
		os.Exit(1)
	}
	if err := app.sdlRenderer.Clear(); err != nil {
		println(err.Error())
		os.Exit(1)
	}
	app.sdlRenderer.Present()
}
