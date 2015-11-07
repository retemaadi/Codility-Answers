package solution

import "testing"

func TestSolution(t *testing.T) {
	expected := 6

	result := Solution(5, []int{1, 3, 1, 4, 2, 3, 5, 4})
	if result != expected {
		t.Error("Expected ", expected, ", Got ", result)
	}
}
