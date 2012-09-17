package ui

import (
	"github.com/Nightgunner5/mmgo/config"
	"github.com/Nightgunner5/mmgo/terrain"
	"github.com/banthar/gl"
)

func drawTerrain() {
	playerX, playerY, _ := player.Position()
	x, y := int64(playerX), int64(playerY)
	x >>= config.ChunkShift - 1
	y >>= config.ChunkShift - 1

	for i := (x - 1) >> 1; i <= (x+1)>>1; i++ {
		for j := (y - 1) >> 1; j <= (y+1)>>1; j++ {
			drawChunk(i, j)
		}
	}
}

func drawChunk(chunkX, chunkY int64) {
	chunk := terrain.GetChunkAt(chunkX, chunkY)

	if chunk.DisplayList == 0 {
		chunk.DisplayList = gl.GenLists(1)

		gl.NewList(chunk.DisplayList, gl.COMPILE)
		for x := 0; x < config.ChunkArraySize() - 1; x++ {
			for y := 0; y < config.ChunkArraySize() - 1; y++ {
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
		gl.EndList()
	}

	gl.CallList(chunk.DisplayList)
}
