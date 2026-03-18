"""Secret Handshake"""

def commands(binary_str):
    """
    Преобразует бинарную строку в список команд согласно правилам:

    Биты (справа налево):
    - бит 0: инвертирует порядок команд (если '1')
    - бит 1: 'jump'
    - бит 2: 'close your eyes'
    - бит 3: 'double blink'
    - бит 4: 'wink'

    Args:
        binary_str (str): Бинарная строка длиной не менее 5 символов.

    Returns:
        list[str]: Список команд в нужном порядке.
    """
    commands_map = [
        ('wink', 4),
        ('double blink', 3),
        ('close your eyes', 2),
        ('jump', 1)
    ]
    result_commands = []
    
    for command, bit_index in commands_map:
        if len(binary_str) > bit_index and binary_str[bit_index] == '1':
            result_commands.append(command)

    if binary_str and binary_str[0] == '1':
        result_commands.reverse()

    return result_commands
