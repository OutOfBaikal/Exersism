def slices(series, length):
    res = []
    s_len = len(series)
    if s_len == 0:
        raise ValueError("series cannot be empty")
    if s_len < length:
        raise ValueError("slice length cannot be greater than series length")
    if length == 0:
        raise ValueError("slice length cannot be zero")
    if length < 0:
        raise ValueError("slice length cannot be negative")
    if s_len == length:
        res.append(series)
        return res
    for index in range(0, s_len - length + 1):
        res.append(series[index:index + length])

    return res
