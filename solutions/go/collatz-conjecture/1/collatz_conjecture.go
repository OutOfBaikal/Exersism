package collatzconjecture

import "fmt"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
        return -1, fmt.Errorf("incorrect value %d", n)
    }
    res := 0
    for n != 1 {
        if n % 2 == 0 {
            n /= 2
        } else {
            n = 3 * n + 1
        }
        res++
    }

    return res, nil
}
