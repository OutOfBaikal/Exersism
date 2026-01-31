package summultiples

func SumMultiples(limit int, divisors ...int) int {
	res := 0
    multiples := make(map[int]struct{})
    for _, divisor := range divisors {
        if divisor == 0 {
            continue // Пропускаем делители, равные нулю
        }
        for i := divisor; i < limit; i += divisor {
            multiples[i] = struct{}{}
        }
    }
    for multiple := range multiples {
        res += multiple
    }
	return res
}
