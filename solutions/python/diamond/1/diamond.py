def rows(letter):
    if not letter.isalpha():
        raise ValueError("Bad symbol!")
    num = (ord(letter) - ord('A')) * 2 + 1
    data = [bytearray(b' ' * num) for _ in range(num)]
    left, right = num // 2, num // 2

    for index in range(0, ord(letter) - ord('A') + 1):
        current_char_code = ord('A') + index
        data[index][left] = current_char_code
        data[index][right] = current_char_code
        left-=1
        right+=1

    for index in range(num // 2 + 1, num):
        sym_index = num - 1 - index
        data[index] = data[sym_index].copy()

    return [row.decode('utf-8') for row in data]
