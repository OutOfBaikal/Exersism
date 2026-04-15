def square_of_sum(number):
    res = 0
    for i in range(1, number + 1):
        res += i
    return res * res


def sum_of_squares(number):
    res = 0
    for i in range(1, number + 1):
        res += i * i
    return res



def difference_of_squares(number):
    return square_of_sum(number) - sum_of_squares(number)
