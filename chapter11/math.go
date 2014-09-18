package math

func Average(xs []float64) float64 {
    total := float64(0)
    for _, x := range xs {
        total += x
    }
    return total / float64(len(xs))
}

func Max(x []float64) float64 {
    biggest := x[0]
    for _, x := range xs {
        if x > biggest {
            biggest = x
        }
    }
    return biggest
}

func Min(x []float64) float64 {
    smallest := x[0]
    for _, x := range xs {
        if x < smallest {
            smallest = x
        }
    }
    return smallest
}

