package vector

import (
	"fmt"
	"math"
)

type Vector []float64

// Returns the first component of this vector. If the vector does not have at
// least one components, this method will panic.
func (v Vector) X() float64 {
	v.assertCapacity(1)
	return v[0]
}

// Returns the second component of this vector. If the vector does not have at
// least two components, this method will panic.
func (v Vector) Y() float64 {
	v.assertCapacity(2)
	return v[1]
}

// Returns the third component of this vector. If the vector does not have at
// least three components, this method will panic.
func (v Vector) Z() float64 {
	v.assertCapacity(3)
	return v[2]
}

// Returns the fourth component of this vector. If the vector does not have at
// least four components, this method will panic.
//
// To access components beyond the fourth, use the Vector as a zero-indexed array.
func (v Vector) W() float64 {
	v.assertCapacity(4)
	return v[3]
}

// Returns a copy of this vector.
func (v Vector) Clone() Vector {
	clone := make(Vector, len(v))

	copy(clone, v)

	return clone
}

// Scales each component of this vector in place by the amount specified.
func (v Vector) Scale(amount float64) Vector {
	for i := range v {
		v[i] *= amount
	}

	return v
}

// Adds another vector to this vector in place. The vectors must be the same degree. Returns this vector.
func (v Vector) Add(other Vector) Vector {
	v.assertSameCapacity(other)

	for i := range v {
		v[i] += other[i]
	}

	return v
}

// Returns the dot product (x1*x2 + y1*y2 + ...) of two vectors. The vectors must be of the same degree.
func (v Vector) Dot(other Vector) float64 {
	v.assertSameCapacity(other)

	var dot float64
	for i := range v {
		dot += v[i] * other[i]
	}

	return dot
}

func (v Vector) String() string {
	if len(v) == 0 {
		return "<Empty>"
	}

	var out string

	out = fmt.Sprintf("<%1.4f", v[0])

	for i, element := range v {
		if i == 0 {
			continue
		}

		out = fmt.Sprintf("%s, %1.4f", out, element)
	}

	return out + ">"
}

func (v Vector) EqualsExactly(other Vector) bool {
	v.assertSameCapacity(other)

	for i := range v {
		if v[i] != other[i] {
			return false
		}
	}

	return true
}

func (v Vector) EqualsApproximately(other Vector, epsilon float64) bool {
	v.assertSameCapacity(other)

	for i := range v {
		if math.Abs(v[i]-other[i]) > epsilon {
			return false
		}
	}

	return true
}

// Panics if the vector is not AT LEAST n-dimensional.
func (v Vector) assertCapacity(n int) {
	if len(v) < n {
		panic(fmt.Sprintf("This operation cannot be used on vectors of less than %d elements.", n))
	}
}

// Panics if this vector does not have the same number of elements as other.
func (v Vector) assertSameCapacity(other Vector) {
	if len(v) != len(other) {
		panic(fmt.Sprintf("This operation is only possible on vectors of the same size. (%d != %d)", len(v), len(other)))
	}
}

// Shortcut vector constructor.
func Vec(components ...float64) Vector {
	return Vector(components)
}

// Constructs a one-dimensional vector. Differs from Vec(x) only in semantics.
func Vec1(x float64) Vector {
	return Vector{x}
}

// Constructs a two-dimensional vector. Differs from Vec(x, y) only in semantics.
func Vec2(x, y float64) Vector {
	return Vector{x, y}
}

// Constructs a three-dimensional vector. Differs from Vec(x, y, z) only in semantics.
func Vec3(x, y, z float64) Vector {
	return Vector{x, y, z}
}

// Constructs a four-dimensional vector. Differs from Vec(x, y, z, w) only in semantics.
func Vec4(x, y, z, w float64) Vector {
	return Vector{x, y, z, w}
}
