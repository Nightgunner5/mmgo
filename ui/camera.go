package ui

import "github.com/banthar/gl"

func positionCamera() {
	gl.LoadIdentity()
	x, y := 0, 0 //player.Position()
	gl.Translatef(float32(-x), float32(-y)-4, -20)
	gl.Rotatef(-30, 1, 0, 0)
}
