def line_up(name, number):
    end = number % 10
    tail = number % 100
    if tail != 11 and tail != 12 and tail != 13: 
        match end:
            case 1:
                return f"{name}, you are the {number}st customer we serve today. Thank you!"
            case 2:
                return f"{name}, you are the {number}nd customer we serve today. Thank you!"
            case 3:
                return f"{name}, you are the {number}rd customer we serve today. Thank you!"
            case _:
                return f"{name}, you are the {number}th customer we serve today. Thank you!"

    return f"{name}, you are the {number}th customer we serve today. Thank you!"