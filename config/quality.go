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
	return terrainQuality>>1 + 2
}

// Do not use this directly. Use the method with the same name in terrain.
func ChangeTerrainQuality(amount int) {
	terrainQuality += amount
	if terrainQuality < 1 {
		terrainQuality = 1
	}
}

func ChunkArraySize() int {
	return ChunkSize*terrainQuality + 1
}
