package solution

// Solution for FrogJmp task (https://codility.com/programmers/task/frog_jmp)
func Solution(X int, Y int, D int) int {
	if X == Y {
		return 0
	}

	s := (Y - X) / D
	if (Y-X)%D == 0 {
		return s
	}

	return s + 1
}
