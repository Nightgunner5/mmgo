package terrain

import "github.com/Nightgunner5/mmgo/config"

type Vec3Array []float64

func (v Vec3Array) Set(x, y int, a, b, c float64) {
	pos := (x*config.ChunkArraySize() + y) * 3
	v[pos], v[pos+1], v[pos+2] = a, b, c
}

func (v Vec3Array) Get(x, y int) (a, b, c float64) {
	pos := (x*config.ChunkArraySize() + y) * 3
	return v[pos], v[pos+1], v[pos+2]
}

type Chunk struct {
	Normals  Vec3Array
	Vertices Vec3Array

	// time.Time.Unix()
	lastGet int64

	DisplayList uint
}
