package resistorcolortrio

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
import (
    "strconv"
    "math"
    )
func Label(colors []string) string {
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
    for i := 0; i < 2; i++ {
        res *= 10
        res += data[colors[i]]
    }
    res = res * int(math.Pow(10.0, float64(data[colors[len(colors) - 1]] + len(colors) - 3)))
    buf := res
    ohms := "ohms"
    a := 0
    for buf >= 1000 {
        buf /= 1000
        a++
    }

    if a == 1 {
        ohms = "kilo" + ohms
    } else if a == 2 {
        ohms = "mega" + ohms
    } else if a == 3 {
        ohms = "giga" + ohms
    }

    return strconv.Itoa(buf) + " " + ohms
}
