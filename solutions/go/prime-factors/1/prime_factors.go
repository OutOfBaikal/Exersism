package prime

import "math"

func Sieve(limit int64) []int64 {
    if limit < 2 {
        return []int64{}
    }
	res := []int64{2}
    n := limit - 1
    size := (n + 1) / 2
    flags := make([]bool, size)
    for i := int64(3); i <= int64(math.Sqrt(float64(limit))); i += 2 {
        if !flags[i / 2] {
            for j := i * i; j <= n; j += i * 2 {
                flags[j / 2] = true
            }
        }
    }
    for i := int64(3); i <= n; i += 2 {
        if !flags[i / 2] {
            res = append(res, i)
        }
    }
    
    return res
}

func Factors(n int64) []int64 {
    if n < 2 {
        return []int64{}
    }
    res := make([]int64, 0)
    var primes []int64
    if n > 1000000000 {
        primes = Sieve(int64(math.Sqrt(float64(n))))
    } else {
		primes = Sieve(n)
    }
    for _, x := range(primes) {
        for n % x == 0 {
            res = append(res, x)
            n /= x
        }
        if n == 1 {
            break
        }
    }
    if n > 1 {
        res = append(res, n)
    }
    return res
}
