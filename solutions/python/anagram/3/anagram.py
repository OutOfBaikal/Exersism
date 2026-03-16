'''solves Anagram exersise'''
from collections import Counter

def find_anagrams(word, candidates):
    '''finds all anagrams for word in candidates'''
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