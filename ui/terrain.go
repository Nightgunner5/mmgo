package ui

import (
	"github.com/Nightgunner5/mmgo/terrain"
	"github.com/banthar/gl"
)

func drawTerrain() {
	var x, y int64 = 0, 0 // player.Position()
	x >>= terrain.ChunkShift
	y >>= terrain.ChunkShift

	for i := x - 2; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			drawChunk(i, j)
		}
	}
}

func drawChunk(chunkX, chunkY int64) {
	chunk := terrain.GetChunkAt(chunkX, chunkY)

	for x := 0; x < terrain.ChunkSizeSubdivisions; x++ {
		for y := 0; y < terrain.ChunkSizeSubdivisions; y++ {
			gl.Normal3d(chunk.Normals.Get(x, y))
			gl.Vertex3d(chunk.Vertices.Get(x, y))
			x++

			gl.Normal3d(chunk.Normals.Get(x, y))
			gl.Vertex3d(chunk.Vertices.Get(x, y))
			y++

			gl.Normal3d(chunk.Normals.Get(x, y))
			gl.Vertex3d(chunk.Vertices.Get(x, y))
			x--

			gl.Normal3d(chunk.Normals.Get(x, y))
			gl.Vertex3d(chunk.Vertices.Get(x, y))
			y--
		}
	}
}
