package perfect

import "errors"

// Define the Classification type here.
type Classification int

const (
    ClassificationDeficient Classification = iota
    ClassificationPerfect
    ClassificationAbundant
)

var ErrOnlyPositive = errors.New("only positive numbers are allowed")

func Classify(n int64) (Classification, error) {
    var c Classification
	if n < 1 {
        return 0, ErrOnlyPositive
    }
    divisors := make([]int64, 0)
    for i := int64(1); i < n / 2 + 1; i++ {
        if n % i == 0 {
            divisors = append(divisors, i)
        }
    }
    sumDiv := sum(divisors)
    if sumDiv == n {
        c = ClassificationPerfect
    } else if sumDiv > n {
        c = ClassificationAbundant
    } else {
        c = ClassificationDeficient
    }
    return c, nil
}

func sum(data []int64) int64 {
    var res int64
    for _, value := range(data) {
        res += value
    }
    return res
}
