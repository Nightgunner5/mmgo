package config

const (
	ChunkShift = 4
	ChunkSize  = 1 << ChunkShift
)

var terrainQuality int = 4

func TerrainSubdivisions() int {
	return terrainQuality
}

func TerrainDetail() int {
	return terrainQuality>>4 + 3
}

func ChunkArraySize() int {
	return ChunkSize*terrainQuality + 1
}
