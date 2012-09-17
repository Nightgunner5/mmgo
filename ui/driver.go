package ui

import (
	"github.com/Nightgunner5/glut"
	"github.com/Nightgunner5/mmgo/terrain"
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
	glut.KeyboardFunc(keydown)
	glut.KeyboardUpFunc(keyup)
	glut.MainLoop()
}

func display() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

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

	gl.Enable(gl.DEPTH_TEST)
	gl.Enable(gl.COLOR_MATERIAL)

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()

	glu.Perspective(45, aspect, NEAR_Z, FAR_Z)

	gl.Enable(gl.LIGHTING)
	gl.Lightfv(gl.LIGHT0, gl.POSITION, []float32{-2, -3, 5, 0})
	gl.Lightfv(gl.LIGHT0, gl.AMBIENT, []float32{0.1, 0.1, 0.2, 1})
	gl.Lightfv(gl.LIGHT0, gl.DIFFUSE, []float32{0.4, 0.35, 0.2, 1})
	gl.Lightfv(gl.LIGHT0, gl.SPECULAR, []float32{1, 1, 1, 1})
	gl.Enable(gl.LIGHT0)

	gl.Rotatef(-30, 1, 0, 0)
	gl.Translatef(0, 0, -5)

	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()

	gl.Lightfv(gl.LIGHT1, gl.AMBIENT, []float32{0, 0, 0, 1})
	gl.Lightfv(gl.LIGHT1, gl.DIFFUSE, []float32{0.4, 0.4, 0.4, 1})
	gl.Lightfv(gl.LIGHT1, gl.SPECULAR, []float32{1, 1, 1, 1})
	gl.Enable(gl.LIGHT1)
}

var previousTime int

func idle() {
	time := glut.Get(glut.ELAPSED_TIME)
	delta := float64(time-previousTime) / 1000
	previousTime = time

	player.Think(delta)

	window.PostRedisplay()
}

// Not synchronized because it is only used in the main loop.
var keys = make(map[byte]bool)

func keydown(key byte, x, y int) {
	keys[key] = true
}

func keyup(key byte, x, y int) {
	delete(keys, key)

	switch key {
	case 'n':
		terrain.ChangeTerrainQuality(-1)
	case 'm':
		terrain.ChangeTerrainQuality(+1)
	}
}

func key(key byte) bool {
	a, b := keys[key]
	return b && a
}

// TODO: put this in a separate package
type player_t struct {
	// Acceleration
	aX, aY, aZ float64

	// Velocity
	vX, vY, vZ float64

	// Position
	x, y, z float64
}

var player player_t

func (p *player_t) Position() (x, y, z float64) {
	return p.x, p.y, p.z
}

func (p *player_t) Think(delta float64) {
	p.aX = 0
	if key('d') {
		p.aX += 1
	}
	if key('a') {
		p.aX -= 1
	}

	if p.vX < 0 {
		p.aX += 0.5
	}
	if p.vX > 0 {
		p.aX -= 0.5
	}

	p.aY = 0
	if key('w') {
		p.aY += 1
	}
	if key('s') {
		p.aY -= 1
	}

	if p.vY < 0 {
		p.aY += 0.5
	}
	if p.vY > 0 {
		p.aY -= 0.5
	}

	p.aZ = -9.8

	p.vX += p.aX * delta
	p.vY += p.aY * delta
	p.vZ += p.aZ * delta

	if p.vX > 5 {
		p.vX = 5
	}
	if p.vX < -5 {
		p.vX = -5
	}

	if p.vY > 5 {
		p.vY = 5
	}
	if p.vY < -5 {
		p.vY = -5
	}

	p.x += p.vX * delta
	p.y += p.vY * delta
	p.z += p.vZ * delta

	terrainHeight := terrain.GetHeightAt(p.x, p.y)
	if p.z < terrainHeight {
		p.z = terrainHeight
		p.vZ = 0
	}
}
