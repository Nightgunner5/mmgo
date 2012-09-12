package vector

import (
	"fmt"
	"math"
)

type Vector []float64

func (v Vector) X() float64 {
	v.AssertCapacity(1)
	return v[0]
}

func (v Vector) Y() float64 {
	v.AssertCapacity(2)
	return v[1]
}

func (v Vector) Z() float64 {
	v.AssertCapacity(3)
	return v[2]
}

func (v Vector) W() float64 {
	v.AssertCapacity(4)
	return v[3]
}

// Scales each component of this vector by amount.
func (v Vector) Scale(amount float64) {
	for i := range v {
		v[i] *= amount
	}
}

func (v Vector) Add(other Vector) {
	v.AssertSameCapacity(other)

	for i := range v {
		v[i] += other[i]
	}
}

func (v Vector) Subtract(other Vector) {
	v.AssertSameCapacity(other)

	for i := range v {
		v[i] -= other[i]
	}
}

func (v Vector) EqualExactly(other Vector) bool {
	v.AssertSameCapacity(other)

	for i := range v {
		if v[i] != other[i] {
			return false
		}
	}

	return true
}

func (v Vector) EqualApproximately(other Vector, epsilon float64) bool {
	v.AssertSameCapacity(other)

	for i := range v {
		if math.Abs(v[i]-other[i]) > epsilon {
			return false
		}
	}

	return true
}

// Panics if the vector is not AT LEAST n-dimensional.
func (v Vector) AssertCapacity(n int) {
	if len(v) < n {
		panic(fmt.Sprintf("This operation cannot be used on vectors of less than %d elements.", n))
	}
}

// Panics if this vector does not have the same number of elements as other.
func (v Vector) AssertSameCapacity(other Vector) {
	if len(v) != len(other) {
		panic(fmt.Sprintf("This operation is only possible on vectors of the same size. (%d != %d)", len(v), len(other)))
	}
}

func Vec1(x float64) Vector {
	return Vector{x}
}

func Vec2(x, y float64) Vector {
	return Vector{x, y}
}

func Vec3(x, y, z float64) Vector {
	return Vector{x, y, z}
}

func Vec4(x, y, z, w float64) Vector {
	return Vector{x, y, z, w}
}
