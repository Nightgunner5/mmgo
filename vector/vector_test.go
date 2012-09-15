package vector

import "testing"

func TestEquality(t *testing.T) {
	vec := Vec(0, 2, 4)
	vec2 := Vec3(0, 2, 4)

	if !vec.EqualsExactly(vec2) {
		t.Errorf("Vector %v != Vector %v (but they should be equal)", vec, vec2)
	}

	vec2 = Vec3(0, 2, 3)

	if vec.EqualsExactly(vec2) {
		t.Errorf("Vector %v == Vector %v (but they shouldn't be equal)", vec, vec2)
	}
}
