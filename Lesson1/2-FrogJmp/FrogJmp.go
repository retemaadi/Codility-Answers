package solution

func Solution(X int, Y int, D int) int {
    if X == Y {
        return 0
    }
    
    var s int = (Y - X) / D
    if (Y - X) % D == 0 {
        return s
    } else {
        return s+1
    }
}
