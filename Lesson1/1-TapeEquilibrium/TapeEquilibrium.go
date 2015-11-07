package solution

func Solution(A []int) int {
    if len(A) == 2 {
        return absDiff(A[0], A[1])
    }
    
    left := A[0]
    right := sum(A) - A[0]
    
    delta := absDiff(left, right)
    min := delta
    for i := 1; i < len(A) - 1; i++ {
        left += A[i]
        right -= A[i]
        delta = absDiff(left, right)
                
        if delta < min {
            min = delta
        }
    }
    
    return min
}

func sum(array []int) int {
    s := 0
    for _,a := range array {
        s += a
    }
    return s
}

func absDiff(a int, b int) int {
    if a > b {
        return a - b
    } else {
        return b - a
    }
}
