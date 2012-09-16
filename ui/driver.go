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
	shift, aspect := 0, float64(width)/float64(height)
	if aspect > 2 {
		aspect = 2
		shift = width/2 - height
		width = height * 2
	}

	gl.Viewport(shift, 0, width, height)
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()

	glu.Perspective(45, aspect, NEAR_Z, FAR_Z)
	gl.Rotatef(-30, 1, 0, 0)
	gl.Translatef(0, 0, -5)

	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()

	gl.Enable(gl.LIGHTING)
	gl.Lightfv(gl.LIGHT0, gl.POSITION, []float32{2, 3, 1, 0})
	gl.Lightfv(gl.LIGHT0, gl.AMBIENT, []float32{0.1, 0.1, 0.2, 1})
	gl.Lightfv(gl.LIGHT0, gl.DIFFUSE, []float32{1, 0.9, 0.6, 1})
	gl.Lightfv(gl.LIGHT0, gl.SPECULAR, []float32{1, 1, 1, 1})
	gl.Enable(gl.LIGHT0)

	gl.Lightfv(gl.LIGHT1, gl.AMBIENT, []float32{0, 0, 0, 1})
	gl.Lightfv(gl.LIGHT1, gl.DIFFUSE, []float32{0.4, 0.4, 0.4, 1})
	gl.Lightfv(gl.LIGHT1, gl.SPECULAR, []float32{1, 1, 1, 1})
	gl.Enable(gl.LIGHT1)
}

var previousTime int

func idle() {
	time := glut.Get(glut.ELAPSED_TIME)
	delta := time - previousTime
	previousTime = time

	player.x += float64(delta) / 1000
	_ = delta

	window.PostRedisplay()
}

type player_t struct{ x, y, z float64 }

var player player_t

func (p *player_t) Position() (x, y, z float64) {
	return p.x, p.y, p.z
}
