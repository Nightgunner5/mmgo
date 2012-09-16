package ui

import "github.com/banthar/gl"

func positionCamera() {
	gl.LoadIdentity()
	x, y, z := player.Position()
	gl.Translated(-x, -y + 4, -z)

	gl.Lightfv(gl.LIGHT1, gl.POSITION, []float32{float32(x), float32(y), float32(z) + 1, 1})
}
