package prime

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
import (
    "fmt"
	"math"
)

func Nth(n int) (int, error) {
    if n < 1 {
        return 0, fmt.Errorf("Invalid input")
    }
    
    // Calculate a more conservative upper limit for the nth prime
    limit := int(float64(n) * (math.Log(float64(n)) + math.Log(math.Log(float64(n)))))
    if limit < 2 {
        limit = 5
    }

    res := make([]int, 0)
    k := limit - 1
    data := make([]int, k)
    for i := 0; i < k; i++ {
        data[i] = i + 2
    }
    
    flags := make([]bool, k)
    for i := 0; i < k; i++ {
        if !flags[i] {
            res = append(res, data[i])
            for j := i + data[i]; j < k; j += data[i] {
                flags[j] = true
            }
        }
    }
    
    if n > len(res) {
        return 0, fmt.Errorf("The nth prime number is out of range")
    }
    
    return res[n - 1], nil
}
