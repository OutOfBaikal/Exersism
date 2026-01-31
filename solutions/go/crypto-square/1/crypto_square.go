package cryptosquare

import (
    "strings"
    "math"
)

func Encode(pt string) string {
	pt = strings.ToLower(pt)
    normalized := ""
    for _, char := range pt {
        if char >= 'a' && char <= 'z' || char >= '0' && char <= '9' { 
        	normalized += string(char)
        }
    }
    n := len(normalized)
    if n == 0 {
        return ""
    }
    
 	r := int(math.Floor(math.Sqrt(float64(n))))
    c := r
    if n > r * c {
        c++
    }
    if c - r > 1 {
        r++
    }
    if c * r < n {
        r++
    }
    grid := make([][]byte, r)
    for i := range grid {
        grid[i] = make([]byte, c)
    }
    index := 0
    for i := 0; i < r; i++ {
        for j := 0; j < c; j++ {
            if index < n {
                grid[i][j] = normalized[index]
                index++
            } else {
                grid[i][j] = ' '
            }
        }
    }

    var encodedWords []string
    for j := 0; j < c; j++ {
        var column strings.Builder
        for i := 0; i < r; i++ {
            column.WriteByte(grid[i][j])
        }
        encodedWords = append(encodedWords, column.String())
    }

    return strings.Join(encodedWords, " ")
}
