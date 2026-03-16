def rebase(input_base, digits, output_base):
    if input_base < 2:
        raise ValueError("input base must be >= 2")
    if output_base < 2:
        raise ValueError("output base must be >= 2")
    num = toDecimal(input_base, digits)
    return fromDecimal(output_base, num)

def toDecimal(input_base, digits):
    res = 0
    n = len(digits)

    for i, x in enumerate(digits):
        if x < 0 or x >= input_base:
            raise ValueError("all digits must satisfy 0 <= d < input base")
        res += x * (input_base ** (n - i - 1))

    return res

def fromDecimal(output_base, num):
    res = []
    if num == 0:
        return [0]
    while num > 0:
        res.append(num % output_base)
        num //= output_base
    n = len(res)

    for i in range(n // 2):
        j = n - i - 1
        res[i], res[j] = res[j], res[i]

    return res