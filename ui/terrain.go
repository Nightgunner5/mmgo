package ui

import (
	"github.com/Nightgunner5/mmgo/terrain"
	"github.com/banthar/gl"
)

func drawTerrain() {
	playerX, playerY, _ := player.Position()
	x, y := int64(playerX), int64(playerY)
	x >>= terrain.ChunkShift - 1
	y >>= terrain.ChunkShift - 1

	for i := (x - 1) >> 1; i <= (x+1)>>1; i++ {
		for j := (y - 1) >> 1; j <= (y+1)>>1; j++ {
			drawChunk(i, j)
		}
	}
}

func tileColor(x, y, z float64) (r, g, b float64) {
	r = float64(int64(x)&3) / 4
	g = float64(int64(y)&3) / 4
	b = z/2 + 0.5
	return
}

func drawChunk(chunkX, chunkY int64) {
	chunk := terrain.GetChunkAt(chunkX, chunkY)

	if chunk.DisplayList == 0 {
		chunk.DisplayList = gl.GenLists(1)

		gl.NewList(chunk.DisplayList, gl.COMPILE)
		for x := 0; x < terrain.ChunkSizeSubdivisions; x++ {
			for y := 0; y < terrain.ChunkSizeSubdivisions; y++ {
				gl.Color3d(tileColor(chunk.Vertices.Get((x/terrain.Subdivisions)*terrain.Subdivisions, (y/terrain.Subdivisions)*terrain.Subdivisions)))

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
