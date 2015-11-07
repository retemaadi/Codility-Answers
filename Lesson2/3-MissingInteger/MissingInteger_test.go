package solution

import "testing"

func TestSolution(t *testing.T) {
	expected := 5

	result := Solution([]int{1, 3, 6, 4, 1, 2})
	if result != expected {
		t.Error("Expected ", expected, ", Got ", result)
	}
}
