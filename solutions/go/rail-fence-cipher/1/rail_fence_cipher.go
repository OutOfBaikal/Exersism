package railfence

func Encode(message string, rails int) string {
	if rails <= 0 {
        return ""
    }
    zigzag := make([]string, rails)
    currentRow, down := 0, true

    for _, char := range message {
        zigzag[currentRow] += string(char)

        if currentRow == 0 {
            down = true
        } else if currentRow == rails - 1 {
            down = false
        }

        if down {
            currentRow++
        } else {
            currentRow--
        }
    }

    result := ""
    for _, row := range zigzag {
        result += row
    }

    return result
}

func Decode(message string, rails int) string {
	if rails <= 0 {
        return ""
    }

    zigzag := make([]string, rails)
    currentRow, down := 0, true

    pos := make([]int, rails)
    for i := 0; i < len(message); i++ {
        pos[currentRow]++
        if currentRow == 0 {
            down = true
        } else if currentRow == rails - 1 {
            down = false
        }
        if down {
            currentRow++
        } else {
            currentRow--
        }
    }

    index := 0
    for r := 0; r < rails; r++ {
        zigzag[r] = message[index: index + pos[r]]
        index += pos[r]
    }

    result := ""
    currentRow, down = 0, true


    for i := 0; i < len(message); i++ {
        result += string(zigzag[currentRow][0])
        zigzag[currentRow] = zigzag[currentRow][1:]

        if currentRow == 0 {
            down = true
        } else if currentRow == rails - 1 {
            down = false
        }

        if down {
            currentRow++
        } else {
            currentRow--
        }
    }

    return result
}
