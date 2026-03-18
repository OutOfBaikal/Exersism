def commands(binary_str):
    res = []
    
    if binary_str[4] == '1':
        res.append("wink")
    if binary_str[3] == '1':
        res.append("double blink")
    if binary_str[2] == '1':
        res.append("close your eyes")
    if binary_str[1] == '1':
        res.append("jump")
    if binary_str[0] == '1':
        res = res[::-1]

    return res
