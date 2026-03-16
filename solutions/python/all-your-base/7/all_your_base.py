"""All Your Base python code"""

def rebase(input_base, digits, output_base):
    """Преобразует числа из одной системы в другую."""
    if input_base < 2:
        raise ValueError("input base must be >= 2")
    if output_base < 2:
        raise ValueError("output base must be >= 2")
    num = to_decimal(input_base, digits)
    return from_decimal(output_base, num)

def to_decimal(input_base, digits):
    """Переводит список цифр в десятичное число."""
    res = 0
    amount = len(digits)

    for index, num in enumerate(digits):
        if num < 0 or num >= input_base:
            raise ValueError("all digits must satisfy 0 <= d < input base")
        res += num * (input_base ** (amount - index - 1))

    return res

def from_decimal(output_base, num):
    """Преобразует десятичное число в список цифр новой базы."""
    res = []
    if num == 0:
        return [0]
    while num > 0:
        res.append(num % output_base)
        num //= output_base
    amount = len(res)

    for index1 in range(amount // 2):
        index2 = amount - index1 - 1
        res[index1], res[index2] = res[index2], res[index1]

    return res