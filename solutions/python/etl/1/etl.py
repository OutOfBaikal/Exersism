def transform(legacy_data):
    result = {}
    for k, v in legacy_data.items():
        for string_item in v:
            result[string_item.lower()] = k

    return result
