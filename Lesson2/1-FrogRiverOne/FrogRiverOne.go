package solution

func Solution(X int, A []int) int {
    e := X*(X+1) / 2
    s := 0
    m := make([]bool, X)
    
    for i,a := range A {
        p := a-1
        if p < X && !m[p] {
            s += a
            if s == e {
                return i
            }
            m[p] = true
        }
    }
    
    return -1
}
