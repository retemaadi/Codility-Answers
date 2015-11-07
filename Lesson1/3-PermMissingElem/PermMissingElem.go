package solution

// Solution for PermMissingElem task (https://codility.com/programmers/task/perm_missing_elem)
func Solution(A []int) int {
	n := len(A)
	if n == 0 {
		return 1
	}

	s := 0
	for _, a := range A {
		s += a
	}

	e := (n + 1) * (n + 2) / 2
	return e - s
}
