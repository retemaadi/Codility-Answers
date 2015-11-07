package solution

import "testing"

func TestSolution(t *testing.T) {
	expected := 1

	result := Solution([]int{3, 1, 2, 4, 3})
	if result != expected {
		t.Error("Expected ", expected, ", Got ", result)
	}
}
