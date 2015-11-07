package solution

func Solution(A []int) int {
    n := len(A)
    if n == 0 {
        return 1
    }
    
    s := 0
    for _,a := range A {
        s += a
    }
    
    e := (n+1)*(n+2) / 2
    return e - s
}
