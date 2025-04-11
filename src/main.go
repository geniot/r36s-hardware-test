// author: HardWareGuy

package main

import (
	"fmt"
	"github.com/go-gl/gl/v3.1/gles2"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	var winTitle string = "Go-SDL2 + Go-GL"
	var winWidth, winHeight int32 = 640, 480
	var window *sdl.Window
	var context sdl.GLContext
	var event sdl.Event
	var running bool
	var err error

	if err = sdl.Init(sdl.INIT_VIDEO); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	if err = gles2.Init(); err != nil {
		panic(err)
	}

	sdl.GLSetAttribute(sdl.GL_CONTEXT_PROFILE_MASK, sdl.GL_CONTEXT_PROFILE_ES)
	//sdl.GLSetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, 3)
	//sdl.GLSetAttribute(sdl.GL_CONTEXT_MINOR_VERSION, 0)
	//sdl.SetHint(sdl.HINT_RENDER_DRIVER, "1")
	sdl.GLSetAttribute(sdl.GL_DOUBLEBUFFER, 1)
	sdl.GLSetAttribute(sdl.GL_DEPTH_SIZE, 24)

	window, err = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight, sdl.WINDOW_OPENGL|sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	//if _, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED); err != nil {
	//	panic(err)
	//}
	context, err = window.GLCreateContext()
	if err != nil {
		panic(err)
	}
	defer sdl.GLDeleteContext(context)

	version, _ := sdl.GetCurrentVideoDriver()
	println(version)

	//println(gles2.GetString(gles2.VERSION))

	//gles2.Disable(gles2.DEPTH_TEST)
	//gles2.ClearColor(1, 0.2, 0.3, 1.0)
	////gles2.ClearDepth(1)
	//gles2.DepthFunc(gles2.LEQUAL)
	//gles2.Viewport(0, 0, int32(winWidth), int32(winHeight))

	running = true
	for running {
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.MouseMotionEvent:
				fmt.Printf("[%d ms] MouseMotion\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n", t.Timestamp, t.Which, t.X, t.Y, t.XRel, t.YRel)
			}
		}
		//gles2.DrawArrays(gles2.TRIANGLES, 0, 0)
		drawgl()
		window.GLSwap()
		sdl.Delay(1)
	}
}

func drawgl() {
	gles2.Clear(gles2.COLOR_BUFFER_BIT)
	gles2.ClearColor(0, 1, 0.3, 1.0)

	//gles2.Begin(gles2.TRIANGLES)
	//gles2.Color3f(1.0, 0.0, 0.0)
	//gles2.Vertex2f(0.5, 0.0)
	//gles2.Color3f(0.0, 1.0, 0.0)
	//gles2.Vertex2f(-0.5, -0.5)
	//gles2.Color3f(0.0, 0.0, 1.0)
	//gles2.Vertex2f(-0.5, 0.5)
	//gles2.End()
}
