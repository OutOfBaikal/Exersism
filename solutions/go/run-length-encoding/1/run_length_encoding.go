package encode

import(
    "strconv"
    "unicode"
)

func RunLengthEncode(input string) string {
	if len(input) == 0 {
        return ""
    }
    var res []byte
    current_letter := rune(input[0])
    current_amount := 1
    for _, x := range input[1:] {
        if x != current_letter {
            if current_amount > 1 {
                res = append(res, []byte(strconv.Itoa(current_amount))...)
            }
            res = append(res, byte(current_letter))
            current_letter = x
            current_amount = 1
        } else {
            current_amount++
        }
    }
    if current_amount > 1 {
        res = append(res, []byte(strconv.Itoa(current_amount))...)
    }
    res = append(res, byte(current_letter))

    return string(res)
}

func RunLengthDecode(input string) string {
    var res []byte
    var data []byte
	for _, x := range input {
		if IsLetter(x) {
            count := 1
            if len(data) > 0 {
                c, err := strconv.Atoi(string(data))
                if err == nil {
                    count = c
                }
                data = data[:0]
            }
            for i := 0; i < count; i++ {
                res = append(res, byte(x))
            }
        } else {
            data = append(data, byte(x))
        }
    }
    return string(res)
}

func IsLetter(a rune) bool {
    return unicode.IsLetter(a) || a == ' '
}
