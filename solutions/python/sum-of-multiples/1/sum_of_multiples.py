def sum_of_multiples(limit, multiples):
    set_res = set()
    res = 0
    for div in multiples:
        if div == 0:
            continue
        for i in range(div, limit, div):
            set_res.add(i)

    for s in set_res:
        res += s
    return res
