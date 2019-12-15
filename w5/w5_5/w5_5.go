package w5_5

func Find(a []int, x int) int {
    switch len(a) {
    case 0:
        return 0
    case 1:
        if x <= a[0] {
            return 0
        }
        return 1
    }