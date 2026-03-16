def compare_dicts(dict1, dict2):
    if len(dict1) != len(dict2):
        return False

    for key, value1 in dict1.items():
        if dict2.get(key) != value1:
            return False

    return True

def find_anagrams(word, candidates):
    result = []
    subject = word.lower()
    from collections import Counter
    pattern = Counter(subject)
    for cand in candidates:
        word = cand.lower()
        if word == subject:
            continue
        data = Counter(word)
        if compare_dicts(pattern, data):
            result.append(cand)

    return result 