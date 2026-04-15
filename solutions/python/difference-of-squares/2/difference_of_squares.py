""" Difference of Squares solution """

def square_of_sum(number):
    res = 0
    for index in range(1, number + 1):
        res += index
    return res * res


def sum_of_squares(number):
    res = 0
    for index in range(1, number + 1):
        res += index * index
    return res



def difference_of_squares(number):
    return square_of_sum(number) - sum_of_squares(number)
