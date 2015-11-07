package solution

import "testing"

func TestSolution(t *testing.T) {
	expected := []int{3, 2, 2, 4, 2}

	result := Solution(5, []int{3, 4, 4, 6, 1, 4, 4})
	if !testEq(expected, result) {
		t.Error("Expected ", expected, ", Got ", result)
	}
}

func testEq(a, b []int) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
