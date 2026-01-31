package lsproduct

import "fmt"

func LargestSeriesProduct(digits string, span int) (int64, error) {
    n := len(digits)
    if span < 0 {
        return 0, fmt.Errorf("span must not be negative")
    }
    if span > n {
        return 0, fmt.Errorf("span must be smaller than string length")
    }
    if n < 3 {
        res := 1
        for i := 0; i < n; i++ {
            res *= int(digits[i] - 48)
        }
        return int64(res), nil
    }
    res := 0
	for i := 0; i < n - span + 1; i++ {
        max := 1
        for j := 0; j < span; j++ {
            if digits[i + j] >= '0' && digits[i + j] <= '9' {
                max *= int(digits[i + j] - 48)
            } else {
                return 0, fmt.Errorf("Invalid value")
            }
        }
        if res < max {
            res = max
        }
    }

    return int64(res), nil
}
