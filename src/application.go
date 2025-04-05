package main

import (
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/tevino/abool/v2"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

type Application struct {
	settings           *Settings
	resources          map[int]*SurfTexture
	sdlWindow          *sdl.Window
	sdlRenderer        *sdl.Renderer
	joysticks          [16]*sdl.Joystick
	pressedKeysCodes   mapset.Set[sdl.Keycode]
	pressedButtonCodes mapset.Set[uint8]
	lastPressedKey     sdl.Keycode
	isRunning          *abool.AtomicBool
	backgroundColor    sdl.Color
}

func NewApplication() *Application {
	return &Application{
		pressedKeysCodes:   mapset.NewSet[sdl.Keycode](),
		pressedButtonCodes: mapset.NewSet[uint8](),
		isRunning:          abool.New(),
		resources:          make(map[int]*SurfTexture),
		backgroundColor:    COLOR_WHITE}
}

func (app *Application) Start(args []string) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		println(err.Error())
		os.Exit(1)
	}
	sdl.JoystickEventState(sdl.ENABLE)
	println("num joysticks: ", sdl.NumJoysticks())

	app.settings = NewSettings()
	app.sdlWindow, _ = sdl.CreateWindow(
		APP_NAME+" "+APP_VERSION,
		int32(app.settings.WindowPosX), int32(app.settings.WindowPosY),
		int32(app.settings.WindowWidth), int32(app.settings.WindowHeight),
		uint32(app.settings.WindowState))
	app.sdlRenderer, _ = sdl.CreateRenderer(app.sdlWindow, -1, sdl.RENDERER_PRESENTVSYNC|sdl.RENDERER_ACCELERATED)
	app.initResources()
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

		case *sdl.JoyAxisEvent:
			// Convert the value to a -1.0 - 1.0 range
			value := float32(t.Value) / 32768.0
			println(value)
			break

		case *sdl.JoyButtonEvent:
			if t.State == sdl.PRESSED {
				app.pressedButtonCodes.Add(t.Button)
			} else {
				app.pressedButtonCodes.Remove(t.Button)
			}
			break

		case *sdl.JoyDeviceAddedEvent:
			// Open joystick for use
			app.joysticks[int(t.Which)] = sdl.JoystickOpen(int(t.Which))
			if app.joysticks[int(t.Which)] != nil {
				fmt.Println("Joystick", t.Which, "connected")
			}
			break
		case *sdl.JoyDeviceRemovedEvent:
			if joystick := app.joysticks[int(t.Which)]; joystick != nil {
				joystick.Close()
			}
			fmt.Println("Joystick", t.Which, "disconnected")
			break

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
	if app.pressedButtonCodes.Contains(BUTTON_CODE_SELECT) && app.pressedButtonCodes.Contains(BUTTON_CODE_START) {
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
	if err := app.sdlRenderer.Copy(app.resources[RESOURCE_BGR_KEY].T, nil, &sdl.Rect{X: 0, Y: 0, W: int32(app.settings.WindowWidth), H: int32(app.settings.WindowHeight)}); err != nil {
		println(err.Error())
	}
	app.sdlRenderer.Present()
}

func (app *Application) initResources() {
	app.resources[RESOURCE_BGR_KEY] = LoadSurfTexture("r36s_blue.png", app.sdlRenderer)
}
