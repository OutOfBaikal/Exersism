def recite(start, take=1):
    mapBottles = ["no", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"]
    end = start - take
    result = []

    for index in range(start, end, -1):
        current = mapBottles[index].capitalize()
        next = mapBottles[index - 1]

        result.append(f"{current} green bottle{pluralize(index)} hanging on the wall,")
        result.append(f"{current} green bottle{pluralize(index)} hanging on the wall,")
        result.append(f"And if one green bottle should accidentally fall,")
        result.append(f"There'll be {next} green bottle{pluralize(index - 1)} hanging on the wall.")
        result.append("")

    return result[:len(result) - 1]

def pluralize(count):
    if count == 1:
        return ""
    return "s"
