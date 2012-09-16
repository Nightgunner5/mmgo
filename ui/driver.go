package ui

import (
	"github.com/Nightgunner5/glut"
	"github.com/banthar/gl"
	"github.com/banthar/glu"
)

const (
	NEAR_Z = 0.1
	FAR_Z  = 1000.0
)

var window glut.Window

func Drive() {
	glut.InitWindowSize(800, 600)
	window = glut.CreateWindow("mmGo")
	positionCamera()
	reshape(800, 600)
	glut.ReshapeFunc(reshape)
	glut.DisplayFunc(display)
	glut.IdleFunc(idle)
	glut.MainLoop()
}

func display() {
	gl.Clear(gl.COLOR_BUFFER_BIT)

	positionCamera()

	gl.Begin(gl.QUADS)
	drawTerrain()
	gl.End()

	gl.Flush()
}

func reshape(width, height int) {
	gl.Viewport(0, 0, width, height)
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()

	aspect := float64(width) / float64(height)
	glu.Perspective(45, aspect, NEAR_Z, FAR_Z)

	gl.Enable(gl.LIGHTING)
	gl.Lightfv(gl.LIGHT0, gl.POSITION, []float32{2, 3, 1, 0})
	gl.Lightfv(gl.LIGHT0, gl.AMBIENT, []float32{0, 0, 0, 1})
	gl.Lightfv(gl.LIGHT0, gl.DIFFUSE, []float32{0.8, 0.8, 0.8, 1})
	gl.Lightfv(gl.LIGHT0, gl.SPECULAR, []float32{1, 1, 0, 1})
	gl.Enable(gl.LIGHT0)

	gl.MatrixMode(gl.MODELVIEW)
}

var previousTime int

func idle() {
	time := glut.Get(glut.ELAPSED_TIME)
	delta := time - previousTime
	previousTime = time

	// TODO: think here
	_ = delta

	window.PostRedisplay()
}
