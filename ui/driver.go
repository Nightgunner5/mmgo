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
	window = glut.CreateWindow("mmGo")
	glut.DisplayFunc(display)
	glut.ReshapeFunc(reshape)
	glut.MainLoop()
}

func display() {
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

	gl.MatrixMode(gl.MODELVIEW)
}
