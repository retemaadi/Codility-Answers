package solution

import "testing"

func TestSolution(t *testing.T) {
	expected := 3

	result := Solution(10, 85, 30)
	if result != expected {
		t.Error("Expected ", expected, ", Got ", result)
	}
}
