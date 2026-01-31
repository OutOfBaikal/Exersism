package hamming

import "fmt"

func Distance(a, b string) (int, error) {
    if len(a) != len(b) {
        return -1, fmt.Errorf("different length")
    }
    res := 0
    for i := 0; i < len(a); i++ {
        if a[i] != b[i] {
            res++
        }
    }
    return res, nil
}
