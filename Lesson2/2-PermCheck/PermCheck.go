package solution

// Solution for PermCheck task (https://codility.com/programmers/task/perm_check)
func Solution(A []int) int {
	n := len(A)
	m := make([]bool, n)

	for _, a := range A {
		p := a - 1
		if p >= n {
			return 0
		}
		if !m[p] {
			m[p] = true
		}
	}

	for _, b := range m {
		if !b {
			return 0
		}
	}
	return 1
}
