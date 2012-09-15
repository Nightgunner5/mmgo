package ui

import "github.com/banthar/gl"

func drawTerrain() {
	x, y := 0, 0 // player.Position()
	x >>= 4
	y >>= 4

	for i := x - 5; i < x+5; i++ {
		for j := y - 1; j < y+2; j++ {
			drawChunk(i, j)
		}
	}
}

func drawChunk(chunkX, chunkY int) {
	color, vertex := func(x, y int) {
		a := float32(Noise(float64(x)/10, float64(y)/10, 0))/4 + 0.5
		gl.Color3f(a, a, a)
	}, func(x, y int) {
		gl.Vertex3f(float32(x), float32(y), float32(Noise(float64(x)/10, float64(y)/10, 0)))
	}

	for x := chunkX << 4; x < (chunkX+1)<<4; x++ {
		for y := chunkY << 4; y < (chunkY+1)<<4; y++ {
			color(x, y)
			vertex(x, y)
			x++

			color(x, y)
			vertex(x, y)
			y++

			color(x, y)
			vertex(x, y)
			x--

			color(x, y)
			vertex(x, y)
			y--
		}
	}
}
