package grains

import (
    "fmt"
    "math"
)

func Square(number int) (uint64, error) {
    if number < 1 || number > 64 {
        return uint64(0), fmt.Errorf("Incorrect value")
    }
	return uint64(math.Pow(2.0, float64(number - 1))), nil
}

func Total() uint64 {
	var res uint64
    res = 0
    for i := 1; i <= 64; i++ {
        data, _ := Square(i)
        res += data
    }

    return res
}
