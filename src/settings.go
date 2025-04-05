package main

import (
	"github.com/magiconair/properties"
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"os"
	"strconv"
)

type Settings struct {
	WindowWidth        int
	WindowHeight       int
	WindowPosX         int
	WindowPosY         int
	WindowState        int
	WindowDisplayIndex int
}

const (
	propDirName           = "/.r36s/"
	propFileName          = "r36s.properties"
	KeyWindowWidth        = "WindowWidth"
	KeyWindowHeight       = "WindowHeight"
	KeyWindowPosX         = "WindowPosX"
	KeyWindowPosY         = "WindowPosY"
	KeyWindowState        = "WindowState"
	KeyWindowDisplayIndex = "WindowDisplayIndex"
)

func (settings *Settings) Save(wnd *sdl.Window) {
	var (
		userHomeDir     string
		propFile        *os.File
		wWidth, wHeight = wnd.GetSize()
		wPosX, wPosY    = wnd.GetPosition()
		displayIndex    int
		props           = properties.NewProperties()
		err             error
	)
	if userHomeDir, err = os.UserHomeDir(); err != nil {
		goto END
	}
	if _, err = os.Stat(userHomeDir + propDirName); os.IsNotExist(err) {
		if err = os.Mkdir(userHomeDir+propDirName, 0777); err != nil {
			goto END
		}
	}
	if propFile, err = os.OpenFile(userHomeDir+propDirName+propFileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777); err != nil {
		goto END
	}
	defer CloseFile(propFile)

	props.MustSet(KeyWindowWidth, strconv.Itoa(int(wWidth)))
	props.MustSet(KeyWindowHeight, strconv.Itoa(int(wHeight)))
	props.MustSet(KeyWindowPosX, strconv.Itoa(int(wPosX)))
	props.MustSet(KeyWindowPosY, strconv.Itoa(int(wPosY)))
	props.MustSet(KeyWindowState, strconv.Itoa(int(wnd.GetFlags())))
	if displayIndex, err = wnd.GetDisplayIndex(); err != nil {
		props.MustSet(KeyWindowDisplayIndex, strconv.Itoa(displayIndex))
	}

	if _, err = props.Write(propFile, properties.UTF8); err != nil {
		goto END
	}
END:
	if err != nil {
		log.Fatal(err)
	}
}

func NewSettings() *Settings {
	var (
		settings         = &Settings{}
		props            = properties.NewProperties()
		userHomeDir      string
		displayMode, err = sdl.GetDesktopDisplayMode(0)
		screenWidth      = int(displayMode.W)
		screenHeight     = int(displayMode.H)
	)
	if userHomeDir, err = os.UserHomeDir(); err == nil {
		if props, err = properties.LoadFile(userHomeDir+propDirName+propFileName, properties.UTF8); err != nil {
			log.Println(err)
			props = properties.NewProperties()
		}
	}
	settings.WindowWidth = 640  //props.GetInt(KeyWindowWidth, If(screenWidth > 800, 800, screenWidth))
	settings.WindowHeight = 480 //props.GetInt(KeyWindowHeight, If(screenHeight > 600, 600, screenHeight))
	settings.WindowPosX = props.GetInt(KeyWindowPosX, (screenWidth-settings.WindowWidth)/2)
	settings.WindowPosY = props.GetInt(KeyWindowPosY, (screenHeight-settings.WindowHeight)/2)
	settings.WindowState = sdl.WINDOW_SHOWN
	settings.WindowDisplayIndex = props.GetInt(KeyWindowDisplayIndex, 0)

	//patching window state
	//settings.WindowState |= sdl.WINDOW_RESIZABLE
	//settings.WindowState |= sdl.WINDOW_MAXIMIZED
	if screenWidth <= 640 && screenHeight <= 480 {
		settings.WindowState |= sdl.WINDOW_BORDERLESS
	}

	return settings
}

func (settings *Settings) SaveWindowState(sdlWindow *sdl.Window) {
	width, height := sdlWindow.GetSize()
	xPos, yPos := sdlWindow.GetPosition()
	settings.WindowState = int(sdlWindow.GetFlags())

	if settings.WindowState&sdl.WINDOW_MAXIMIZED <= 0 {
		settings.WindowWidth = int(width)
		settings.WindowHeight = int(height)
		settings.WindowPosX = int(xPos)
		settings.WindowPosY = int(yPos)
	}
}
