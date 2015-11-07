package solution

// Solution for MissingInteger task (https://codility.com/programmers/task/missing_integer)
func Solution(A []int) int {
	const MAX = 100001
	m := make([]bool, MAX)

	for _, a := range A {
		p := a - 1
		if p < 0 || p >= MAX {
			continue
		}
		if !m[p] {
			m[p] = true
		}
	}

	for i, b := range m {
		if !b {
			return i + 1
		}
	}
	return MAX
}
