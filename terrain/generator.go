package terrain

import "github.com/Nightgunner5/mmgo/vector"

const (
	scaleHorizontal = 0.4
	scaleVertical   = 1.5

	noiseIterations = 3
	noiseAmplitude  = 0.15
	noiseFrequency  = 3
)

func terrainHeightExpensive(x, y float64) float64 {
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

func generateChunk(chunkCoord ChunkCoordinate) *Chunk {
	startX, startY := float64(chunkCoord.X<<ChunkShift), float64(chunkCoord.Y<<ChunkShift)

	coord := func(i, j int) (x, y float64) {
		return startX + float64(i)*SubdivisionSize, startY + float64(j)*SubdivisionSize
	}

	// Optimization: Pre-generate the noise values we'll be using instead of generating them repeatedly
	var heights [ChunkSizeSubdivisions + Subdivisions + 1][ChunkSizeSubdivisions + Subdivisions + 1]float64
	for i := range heights {
		for j := range heights[i] {
			heights[i][j] = terrainHeightExpensive(coord(i, j))
		}
	}

	chunk := new(Chunk)

	for i := 0; i < ChunkSizeSubdivisions+1; i++ {
		for j := 0; j < ChunkSizeSubdivisions+1; j++ {
			x, y := coord(i, j)

			// Normal
			current, right, top := heights[i][j], heights[i+Subdivisions][j], heights[i][j+Subdivisions]
			v := vector.Vec3(current-right, current-top, scaleVertical / 2)
			v.Normalize()
			chunk.Normals.Set(i, j, v[0], v[1], v[2])

			// Position
			chunk.Vertices.Set(i, j, x, y, current)
		}
	}

	chunk.markGet()

	return chunk
}
