package eliudseggs

func EggCount(displayValue int) int {
	var res int
    for i := 0; i < 32; i++ {
        res += ((displayValue >> i) & 1)
    }

    return res
}
