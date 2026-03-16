from collections import Counter

def find_anagrams(word, candidates):
    result = []
    subject = word.lower()
    pattern = Counter(subject)
    
    for cand in candidates:
        word = cand.lower()
        if word == subject:
            continue
        if Counter(word) == pattern:
            result.append(cand)

    return result 