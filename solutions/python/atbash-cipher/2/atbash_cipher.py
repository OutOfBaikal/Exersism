"""Atbash cipher implementation"""
ENC_ELEMENTS = 'zyxwvutsrqponmlkjihgfedcba'

def encode(plain_text):
    """atbash encoding"""
    encoded = ''
    count = 0
    for symbol in plain_text:
        count += 1
        if symbol.isalpha():
            ascii_code = ord(symbol)
            if ord('A') <= ascii_code <= ord('Z'):
                encoded += ENC_ELEMENTS[ascii_code - ord('A')]
            else:
                encoded += ENC_ELEMENTS[ascii_code - ord('a')]
        elif symbol.isdigit():
            encoded += symbol
        else:
            count -= 1
        if count == 5:
            encoded += ' '
            count = 0
    if encoded[-1] == ' ':
        return encoded[:-1]
    return encoded


def decode(ciphered_text):
    """atbash decoding"""
    decoded = ''
    for symbol in ciphered_text:
        if symbol.isalpha():
            ascii_code = ord(symbol)
            if ord('A') <= ascii_code <= ord('Z'):
                decoded += ENC_ELEMENTS[ascii_code - ord('A')]
            else:
                decoded += ENC_ELEMENTS[ascii_code - ord('a')]
        elif symbol.isdigit():
            decoded += symbol
    return decoded
