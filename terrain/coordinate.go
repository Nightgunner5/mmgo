package terrain

import "sync"

type ChunkCoordinate struct {
	X, Y int64
}

var chunks = make(map[ChunkCoordinate]*Chunk)
var chunkLock sync.RWMutex

func GetChunk(coord ChunkCoordinate) *Chunk {
	chunkLock.RLock()
	if chunk, exists := chunks[coord]; exists {
		chunkLock.RUnlock()
		chunk.markGet()
		return chunk
	}
	chunkLock.RUnlock()

	chunkLock.Lock()
	defer chunkLock.Unlock()
	// Avoid race condition
	if chunk, exists := chunks[coord]; exists {
		chunk.markGet()
		return chunk
	}

	chunks[coord] = generateChunk(coord)
	return chunks[coord]
}

func GetChunkAt(x, y int64) *Chunk {
	return GetChunk(ChunkCoordinate{x, y})
}
