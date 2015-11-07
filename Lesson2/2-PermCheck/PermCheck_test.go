package solution

import "testing"

func TestSolution_1(t *testing.T) {
	expected := 1

	result := Solution([]int{4, 1, 3, 2})
	if result != expected {
		t.Error("Expected ", expected, ", Got ", result)
	}
}

func TestSolution_2(t *testing.T) {
	expected := 0

	result := Solution([]int{4, 1, 3})
	if result != expected {
		t.Error("Expected ", expected, ", Got ", result)
	}
}
