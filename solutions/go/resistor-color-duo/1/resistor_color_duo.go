package resistorcolorduo

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
    res := 0
    data := make(map[string]int, 0)
    data["black"] = 0
    data["brown"] = 1
    data["red"] = 2
    data["orange"] = 3
    data["yellow"] = 4
    data["green"] = 5
    data["blue"] = 6
    data["violet"] = 7
    data["grey"] = 8
    data["white"] = 9
	for i := 0; i < len(colors); i++ {
        res *= 10
        res += data[colors[i]]
    }
    if res > 99 {
        res /= 10
    }

    return res
}
