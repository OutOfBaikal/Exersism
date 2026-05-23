"""Eliud's Eggs solution"""
def egg_count(display_value):
    "function"
    res = 0
    for index in range(0, 32):
        res += ((display_value >> index) & 1)

    return res
