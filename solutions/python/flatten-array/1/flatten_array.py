def flatten(iterable, result = None):
    if result is None:
        result = []
    for item in iterable:
        if isinstance(item, list):
            flatten(item, result)
        elif item != None:
            result.append(item)
    return result
