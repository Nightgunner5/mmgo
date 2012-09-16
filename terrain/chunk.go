package terrain

const (
	Subdivisions    = 4
	SubdivisionSize = 1.0 / Subdivisions

	ChunkShift            = 4
	ChunkSize             = 1 << ChunkShift
	ChunkSizeSubdivisions = ChunkSize * Subdivisions
)

type Vec3Array [(ChunkSizeSubdivisions+1) * (ChunkSizeSubdivisions+1) * 3]float64

func (v *Vec3Array) Set(x, y int, a, b, c float64) {
	pos := (x*(ChunkSizeSubdivisions+1) + y) * 3
	v[pos], v[pos+1], v[pos+2] = a, b, c
}

func (v *Vec3Array) Get(x, y int) (a, b, c float64) {
	pos := (x*(ChunkSizeSubdivisions+1) + y) * 3
	return v[pos], v[pos+1], v[pos+2]
}

type Chunk struct {
	Normals  Vec3Array
	Vertices Vec3Array
}
