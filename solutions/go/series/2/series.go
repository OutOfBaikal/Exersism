package series

func All(n int, s string) []string {
	res := make([]string, 0)
    if len(s) < n || n <= 0 {
        return res
    }
    if len(s) == n {
        res = append(res, s)
        return res
    }
    for i := 0; i < len(s) - n + 1; i++ {
        res = append(res, s[i:i + n])
    }
    return res
}

func UnsafeFirst(n int, s string) string {
    if len(s) <= n {
        return s
    }
    return s[0:n]
}
