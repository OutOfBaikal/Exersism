package resistorcolor

// Colors returns the list of all colors.
func Colors() []string {
    data := []string{"black", "brown", "red",
			"orange", "yellow", "green", "blue",
			"violet", "grey", "white"}
	return data
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
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
    return data[color]
}
