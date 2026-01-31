package rotationalcipher

func RotationalCipher(plain string, shiftKey int) string {
	data := []rune(plain)
    for i := 0; i < len(data); i++ {
        if data[i] >= 'A' && data[i] <= 'Z' {
        	data[i] = (data[i] - 'A' + rune(shiftKey)) % 26 + 'A'
        }
        if data[i] >= 'a' && data[i] <= 'z' {
        	data[i] = (data[i] - 'a' + rune(shiftKey)) % 26 + 'a'
        }
    }

    return string(data)
}
