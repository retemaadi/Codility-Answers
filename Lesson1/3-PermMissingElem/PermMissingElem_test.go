package solution

import "testing"

func TestSolution(t *testing.T) {
	expected := 4

	result := Solution([]int{2, 3, 1, 5})
	if result != expected {
		t.Error("Expected ", expected, ", Got ", result)
	}
}
