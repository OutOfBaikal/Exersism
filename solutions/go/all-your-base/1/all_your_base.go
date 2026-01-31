package allyourbase

import (
    "errors"
    "math"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
        return nil, errors.New("input base must be >= 2")
    }
    if outputBase < 2 {
        return nil, errors.New("output base must be >= 2")
    }
    num, err := ToDecimal(inputBase, inputDigits)
    if err != nil {
        return nil, err
    }
    return FromDecimal(outputBase, num), nil
}

func ToDecimal(inputBase int, inputDigits []int) (int, error) {
    res := 0
    n := len(inputDigits)
    
    for i, x := range inputDigits {
        if x < 0 || x >= inputBase {
            return 0, errors.New("all digits must satisfy 0 <= d < input base")
        }
        res += x * int(math.Pow(float64(inputBase), float64(n - i - 1)))
    }
    return res, nil
}

func FromDecimal(outputBase int, inputDigits int) []int {
    res := []int{}
    if inputDigits == 0 {
        return []int{0}
    }
    for inputDigits > 0 {
        res = append(res, inputDigits % outputBase)
        inputDigits /= outputBase
    }
    n := len(res)
    for i := 0; i < n / 2; i++ {
        j := n - i - 1
        res[i], res[j] = res[j], res[i]
    }
    return res
}