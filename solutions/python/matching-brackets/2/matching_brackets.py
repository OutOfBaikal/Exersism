"""Matching brackets"""
from collections import deque

def is_paired(input_string):
    """Проверяем, все ли скобки нашли свою пару... мяу! 🐾"""
    stack = deque()
    pairs = { ']': '[', '}': '{', ')': '('}
    for symbol in input_string:
        match symbol:
            case '[' | '{' | '(':
                stack.append(symbol)
            case ']' | '}' | ')':
                if not stack:
                    return False
                if pairs[symbol] != stack.pop():
                    return False

    return len(stack) == 0