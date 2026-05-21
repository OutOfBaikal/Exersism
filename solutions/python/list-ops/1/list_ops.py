def append(list1, list2):
    result = list1.copy()
    for x in list2:
        result.append(x)
    return result


def concat(lists):
    result = []
    for item in lists:
        if isinstance(item, list):
            for sub_item in item:
                result.append(sub_item)
        elif item != None:
            result.append(item)
    return result


def filter(function, lst):
    result = []
    for x in lst:
        if function(x):
            result.append(x)
    return result


def length(lst):
    return len(lst)


def map(function, lst):
    result = []
    for x in lst:
        result.append(function(x))
    return result


def foldl(function, lst, initial):
    for x in lst:
        initial = function(initial, x)
    return initial


def foldr(function, lst, initial):
    for x in lst[::-1]:
        initial = function(initial, x)
    return initial



def reverse(lst):
    return lst[::-1]
