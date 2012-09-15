package ui

import (
	"github.com/banthar/gl"

//	"github.com/banthar/glu"
)

func DisplayMenu() {
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.Begin(gl.QUADS)
	gl.Color3f(1, 1, 1)
	gl.Vertex2f(-0.6, -0.8)
	gl.Vertex2f(+0.6, -0.8)
	gl.Vertex2f(+0.6, +0.8)
	gl.Vertex2f(-0.6, +0.8)
	gl.End()
	gl.Flush()
}

func ReshapeMenu(width, height int) {
	gl.Viewport(0, 0, width, height)
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	aspect := float64(width) / float64(height)
	w, h := 1.0, 1.0
	if aspect > 1 {
		w = aspect
	} else {
		h = 1 / aspect
	}
	gl.Ortho(-w, w, -h, h, -1, 1)
}
