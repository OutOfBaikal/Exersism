package diamond

import (
    "bytes"
    "errors"
)

func Gen(char byte) (string, error) {
    if char > 'Z' || char < 'A' {
        return "", errors.New("Bad symbol!")
    }
	n := (int(char) - int('A')) * 2 + 1
    data := make([][]byte, n)
    for i := range data {
        data[i] = make([]byte, n)
        for j := range data[i] {
            data[i][j] = ' '
        }
    }
    left, right := n / 2, n / 2
    
    for i := 0; i <= int(char-'A'); i++ {
        data[i][left] = byte('A' + i)
        data[i][right] = byte('A' + i)
        left--
        right++
    }
    
    left = 1
    right = n - 2
    for i := n / 2 + 1; i < n; i++ {
        symIndex := n - 1 - i
        data[i][left] = data[symIndex][left]
        data[i][right] = data[symIndex][right]
        left++
        right--
    }

    result := string(bytes.Join(data, []byte{'\n'}))
    return result, nil
}
