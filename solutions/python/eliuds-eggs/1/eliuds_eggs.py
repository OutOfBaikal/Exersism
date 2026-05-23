def egg_count(display_value):
    res = 0
    for i in range(0, 32):
        res += ((display_value >> i) & 1)

    return res
