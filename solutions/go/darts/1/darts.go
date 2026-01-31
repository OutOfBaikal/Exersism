package darts

import "math"

func Score(x, y float64) int {
    score := math.Sqrt(x * x + y * y)
	if score > 10.0 {
        return 0
    } else if score <= 10 && score > 5 {
        return 1
    } else if score <= 5 && score > 1 {
        return 5
    }
    return 10
}
