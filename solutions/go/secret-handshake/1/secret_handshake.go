package secrethandshake

func Handshake(code uint) []string {
    commands_map := []string{"wink", "double blink", "close your eyes", "jump", "reverse"}
    var result []string
    i := 0
    for code > 0 {
        if code & 1 == 1 {
            if commands_map[i] != "reverse" {
            	result = append(result, commands_map[i])
            } else {
                for j := 0; j < len(result) / 2; j++ {
                    result[j], result[len(result) - j - 1] = result[len(result) - j - 1], result[j]
                }
            }
        }
        i += 1
        code >>= 1
    }

    return result
}
