package main

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"

	"fmt"
	"os"
	"runtime"
	"strings"
)

const (
	SCR_WIDTH  = 640
	SCR_HEIGHT = 480
)

const vertexShaderSource = `#version 330 core
layout (location = 0) in vec3 aPos;
void main() {
	gl_Position = vec4(aPos.xyz, 1.0);
}
` + "\x00"

const fragmentShaderSource = `#version 330 core
out vec4 FragColor;
void main() {
	FragColor = vec4(1.0, 0.5, 0.2, 1.0);
}
` + "\x00"

var (
	aPos uint32 = 0 // by layout (location = 0)
)

func init() {
	runtime.LockOSThread()
}

func main() {

	// glfw: initialize and configure
	glfw.Init()
	defer glfw.Terminate()
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	// glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True) // apple only

	// glfw: window creation
	window, err := glfw.CreateWindow(SCR_WIDTH, SCR_HEIGHT, "Hello triangle", nil, nil)
	if err != nil {
		fail("create GLFW window", err, -1)
		return
	}
	window.MakeContextCurrent()
	window.SetFramebufferSizeCallback(onResize)

	// gl: initialize
	if err := gl.Init(); err != nil {
		fail("init GL", err, -1)
		return
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)

	// create shader program
	program := createProgram(vertexShaderSource, fragmentShaderSource)
	_ = program

	// set up vetex data (and buffer(s)) and configure vertex attributes

	vertices := []float32{
		0.5, 0.5, 0.0, // top right
		0.5, -0.5, 0.0, // bottom right
		-0.5, -0.5, 0.0, // bottom left
		-0.5, 0.5, 0.0, // top left
	}
	indices := []uint32{
		0, 1, 3, // first Triangle
		1, 2, 3, // second Triangle
	}

	var (
		VAO uint32
		VBO uint32
		EBO uint32
	)
	// VAO {
	// 	EBO
	// }

	gl.GenVertexArrays(1, &VAO)
	gl.GenBuffers(1, &VBO)
	gl.GenBuffers(1, &EBO)
	// bind the Vertex Array Object first,
	// then bind and set vetex buffer(s),
	// and then configure vertex attributes(s).
	gl.BindVertexArray(VAO)

	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, EBO)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)

	gl.VertexAttribPointer(aPos, 3, gl.FLOAT, false, 3*4, gl.Ptr(nil))
	gl.EnableVertexAttribArray(aPos)

	// note that this is allowed,
	// the call to glVertexAttribPointer registered VBO
	// as the vertex attribute's bound vertex buffer object
	// so afterwards we can safely unbind
	gl.BindBuffer(gl.ARRAY_BUFFER, 0) // VBO

	// remember: do NOT unbind the EBO while a VAO is active
	// as the bound element buffer object IS stored in the VAO;
	// keep the EBO bound.
	// !!! glBindBuffer(GL_ELEMENT_ARRAY_BUFFER, 0); // EBO

	// You can unbind the VAO afterwards
	// so other VAO calls won't accidentally modify this VAO,
	// but this rarely happens.
	// Modifying other VAOs requires a call to glBindVertexArray anyways
	// so we generally don't unbind VAOs (nor VBOs) when it's not directly necessary.
	gl.BindVertexArray(0) // VAO

	// render loop
	for !window.ShouldClose() {

		// render
		gl.ClearColor(0.2, 0.3, 0.3, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		// draw out first triangle
		gl.UseProgram(program)
		gl.BindVertexArray(VAO)
		// seeing as we only have a single VAO there's no need to bind
		// it every time, but we'll do so to keep things a bit more organized
		gl.DrawElements(gl.TRIANGLES, 3, gl.UNSIGNED_INT, gl.Ptr(nil))

		// swap buffers and poll IO events
		window.SwapBuffers()
		glfw.PollEvents()

	}

	gl.DeleteVertexArrays(1, &VAO)
	gl.DeleteBuffers(1, &VBO)
	gl.DeleteBuffers(1, &EBO)

}

func createProgram(vertexShaderSource, fragmentShaderSource string) uint32 {
	// compile shaders
	vertexShader := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	fragmentShader := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)

	// link shaders
	program := gl.CreateProgram()
	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	//infoLog := getProgramLog(program)
	//var status int32
	//if gl.GetProgramiv(program, gl.LINK_STATUS, &status); status == gl.FALSE {
	//	fail("link program "+infoLog, nil, -3)
	//} else {
	//	fmt.Printf("log of program:\n%s\n", infoLog)
	//}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return program
}

func compileShader(shaderSrc string, shaderType uint32) uint32 {
	csources, free := gl.Strs(shaderSrc)
	defer free()
	shader := gl.CreateShader(shaderType)
	gl.ShaderSource(shader, 1, csources, nil)
	gl.CompileShader(shader)

	// check for compile errors
	//infoLog := getShaderLog(shader)
	//var status int32
	//if gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status); status == gl.FALSE {
	//	fail("compile "+shaderTypeToString(shaderType)+" shader "+infoLog, nil, -2)
	//} else {
	//	fmt.Printf("log of %s shader:\n%s\n", shaderTypeToString(shaderType), infoLog)
	//}
	return shader
}

func getShaderLog(shader uint32) string {
	var logLen int32
	gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLen)
	infoLog := strings.Repeat("\x00", int(logLen))
	gl.GetShaderInfoLog(shader, logLen, nil, gl.Str(infoLog))
	return infoLog
}

func getProgramLog(program uint32) string {
	var logLen int32
	gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLen)
	infoLog := strings.Repeat("\x00", int(logLen))
	gl.GetProgramInfoLog(program, logLen, nil, gl.Str(infoLog))
	return infoLog
}

func shaderTypeToString(shaderType uint32) string {
	shaderTypeStr := (map[uint32]string{
		gl.VERTEX_SHADER:   "vertex",
		gl.FRAGMENT_SHADER: "fragment",
	})[shaderType]
	return shaderTypeStr
}

func onResize(window *glfw.Window, w, h int) {

}

func fail(msg string, err error, code int) {
	fmt.Fprintf(os.Stderr, "Failed to %s: %v\n", msg, err)
	os.Exit(code)
}
