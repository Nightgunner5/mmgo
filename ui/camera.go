package ui

import "github.com/banthar/gl"

func positionCamera() {
	gl.LoadIdentity()
	x, y := 0, 0 //player.Position()
	gl.Translatef(float32(-x), float32(-y), 0)
	gl.Rotatef(-30, 1, 0, 0)
	gl.Translatef(0, 0, -5)
}
