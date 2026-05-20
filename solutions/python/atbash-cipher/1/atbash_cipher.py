enc_elements = 'zyxwvutsrqponmlkjihgfedcba'

def encode(plain_text):
    encoded = ''
    count = 0
    for x in plain_text:
        count += 1
        if x.isalpha():
            a = ord(x)
            if a >= ord('A') and a <= ord('Z'):
                encoded += enc_elements[a - ord('A')]
            else:
                encoded += enc_elements[a - ord('a')]
        elif x.isdigit():
            encoded += x
        else:
            count -= 1
        if count == 5:
            encoded += ' '
            count = 0
    if encoded[-1] == ' ':
        return encoded[:-1]
    return encoded


def decode(ciphered_text):
    decoded = ''
    for x in ciphered_text:
        if x.isalpha():
            a = ord(x)
            if a >= ord('A') and a <= ord('Z'):
                decoded += enc_elements[a - ord('A')]
            else:
                decoded += enc_elements[a - ord('a')]
        elif x.isdigit():
            decoded += x
    return decoded
