package squareroot

import (
    "math"
    "errors"
)

func SquareRoot(number int) (int, error) {
    if number < 0 {
        return -1, errors.New("Negative value")
    }
	num := float32(number)
    x2 := num * 0.5
    threehalfs := float32(1.5)

    i := math.Float32bits(num)
    i = 0x5f3759df - (i >> 1)
    f := math.Float32frombits(i)
    f = f * (threehalfs - (x2 * f * f))

    return int(1.0 / f), nil
}
