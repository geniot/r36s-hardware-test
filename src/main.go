// author: HardWareGuy

package main

import (
	"fmt"
	"github.com/go-gl/gl/v3.2-core/gl"
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

	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	//sdl.GLSetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, 3)
	//sdl.GLSetAttribute(sdl.GL_CONTEXT_MINOR_VERSION, 0)

	window, err = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight, sdl.WINDOW_OPENGL)
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

	//version := gl.GoStr(gl.GetString(gl.VERSION))
	//println(version)

	//println(gl.GetString(gl.VERSION))

	if err = gl.Init(); err != nil {
		panic(err)
	}

	//gl.Disable(gl.DEPTH_TEST)
	//gl.ClearColor(1, 0.2, 0.3, 1.0)
	////gl.ClearDepth(1)
	//gl.DepthFunc(gl.LEQUAL)
	//gl.Viewport(0, 0, int32(winWidth), int32(winHeight))

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
		//gl.DrawArrays(gl.TRIANGLES, 0, 0)
		drawgl()
		window.GLSwap()
		sdl.Delay(1)
	}
}

func drawgl() {
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.ClearColor(0, 1, 0.3, 1.0)

	//gl.Begin(gl.TRIANGLES)
	//gl.Color3f(1.0, 0.0, 0.0)
	//gl.Vertex2f(0.5, 0.0)
	//gl.Color3f(0.0, 1.0, 0.0)
	//gl.Vertex2f(-0.5, -0.5)
	//gl.Color3f(0.0, 0.0, 1.0)
	//gl.Vertex2f(-0.5, 0.5)
	//gl.End()
}
