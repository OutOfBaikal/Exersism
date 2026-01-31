package thefarm

import (
    "fmt"
    "errors"
)

// TODO: define the 'DivideFood' function
func DivideFood(calc FodderCalculator, num int) (float64, error) {
    fodder, err1 := calc.FodderAmount(num)
	factor, err2 := calc.FatteningFactor()
    if err1 == nil && err2 == nil && num > 0 {
        return fodder * factor / float64(num), nil
    } 
	if err1 != nil {
        return 0, err1
    }
	return 0, err2
}
// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(calc FodderCalculator, num int) (float64, error) {
    if num <= 0 {
        return 0, errors.New("invalid number of cows")
    }
    fodder, err1 := calc.FodderAmount(num)
	factor, err2 := calc.FatteningFactor()
    if err1 == nil && err2 == nil {
        return fodder * factor / float64(num), nil
    } 
    if err1 != nil {
        return 0, err1
    }
	return 0, err2
}
// TODO: define the 'ValidateNumberOfCows' function
type InvalidCowsError struct {
	num int
    details string
}
func (e *InvalidCowsError) Error() string {
    return fmt.Sprintf("%d cows are invalid: %s", e.num, e.details)
}
func ValidateNumberOfCows(num int) error {
    if num < 0 {
        return &InvalidCowsError{
            num: num,
            details: "there are no negative cows",
        }
    } else if num == 0 {
        return &InvalidCowsError{
            num: 0,
            details: "no cows don't need food",
        }
    }
    return nil
}


// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
