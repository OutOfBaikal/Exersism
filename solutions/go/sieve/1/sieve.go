package sieve

func Sieve(limit int) []int {
	res := make([]int, 0)
    n := limit - 1
    data := make([]int, n)
    for i := 0; i < n; i++ {
		data[i] = i + 2
	}
    flags := make([]bool, n)

    for i := 0; i < n; i++ {
        if !flags[i] {
            res = append(res, data[i])
            for j := i + data[i]; j < n; j += data[i] {
                flags[j] = true
            }
        }
    }
    
    return res
}
