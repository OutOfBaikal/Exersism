def pig_latin(word):
    vowels = "aeiou"
    if word.startswith(('xr', 'yt')) or word[0] in vowels:
        return word + 'ay'

    for i in range(len(word) - 1):
        if word[i:i + 2] == 'qu':
            if not any(char in vowels for char in word[:i]):
                return word[i + 2:] + word[:i + 2] + 'ay'
            break

    for i in range(1, len(word)):
        if word[i] == 'y':
            if not any(char in vowels for char in word[:i]):
                return word[i:] + word[:i] + 'ay'
            break

    for i, letter in enumerate(word):
        if letter in vowels:
            return word[i:] + word[:i] + 'ay'

    return word + 'ay'

def translate(text):
    words = text.split()
    pig_latin_words = [pig_latin(word) for word in words]
    return ' '.join(pig_latin_words)
    