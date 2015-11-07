package solution

// Solution for MaxCounters task (https://codility.com/programmers/task/max_counters)
func Solution(N int, A []int) []int {
	c := make([]int, N)

	max, lastMax := 0, 0
	for _, a := range A {
		if a == N+1 {
			lastMax = max
		} else {
			p := a - 1
			if c[p] < lastMax {
				c[p] = lastMax + 1
			} else {
				c[p]++
			}

			if c[p] > max {
				max = c[p]
			}
		}
	}

	for j, m := range c {
		if m < lastMax {
			c[j] = lastMax
		}
	}
	return c
}
