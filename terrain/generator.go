package terrain

import (
	"github.com/Nightgunner5/mmgo/config"
	"github.com/Nightgunner5/mmgo/vector"
)

const (
	scaleHorizontal = 0.4
	scaleVertical   = 1.5

	traceNoiseIterations = 3
	noiseAmplitude       = 0.3
	noiseFrequency       = 3
)

func terrainHeight(x, y float64, noiseIterations int) float64 {
	value, max, mul := 0.0, 0.0, 1.0

	for i := 0; i < noiseIterations; i++ {
		n := Noise(x*scaleHorizontal, y*scaleHorizontal, 0) * scaleVertical

		value += n * mul
		max += mul

		x, y = x*noiseFrequency, y*noiseFrequency
		mul *= noiseAmplitude
	}

	return value / max
}

func GetHeightAt(x, y float64) (z float64) {
	return terrainHeight(x, y, traceNoiseIterations)
}

func generateChunk(chunkCoord ChunkCoordinate) *Chunk {
	startX, startY := float64(chunkCoord.X<<config.ChunkShift), float64(chunkCoord.Y<<config.ChunkShift)

	coord := func(i, j int) (x, y float64) {
		return startX + float64(i)/float64(config.TerrainSubdivisions()), startY + float64(j)/float64(config.TerrainSubdivisions())
	}

	// Optimization: Pre-generate the noise values we'll be using instead of generating them repeatedly
	heightCount := config.ChunkArraySize() + config.TerrainSubdivisions()
	heights := make([][]float64, heightCount)
	for i := range heights {
		heights[i] = make([]float64, heightCount)
		for j := range heights[i] {
			x, y := coord(i, j)
			heights[i][j] = terrainHeight(x, y, config.TerrainDetail())
		}
	}

	chunk := new(Chunk)
	chunk.Vertices = make(Vec3Array, config.ChunkArraySize()*config.ChunkArraySize()*3)
	chunk.Normals = make(Vec3Array, config.ChunkArraySize()*config.ChunkArraySize()*3)

	for i := 0; i < config.ChunkArraySize(); i++ {
		for j := 0; j < config.ChunkArraySize(); j++ {
			x, y := coord(i, j)

			// Normal
			current, right, top := heights[i][j], heights[i+config.TerrainSubdivisions()][j], heights[i][j+config.TerrainSubdivisions()]
			v := vector.Vec3(current-right, current-top, scaleVertical/2)
			v.Normalize()
			chunk.Normals.Set(i, j, v[0], v[1], v[2])

			// Position
			chunk.Vertices.Set(i, j, x, y, current)
		}
	}

	chunk.markGet()

	return chunk
}
