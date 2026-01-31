package twelve

import "fmt"

func Verse(i int) string {
	numbers := make(map[int]string)
    numbers[1] = "first"
    numbers[2] = "second"
    numbers[3] = "third"
    numbers[4] = "fourth"
    numbers[5] = "fifth"
    numbers[6] = "sixth"
    numbers[7] = "seventh"
    numbers[8] = "eighth"
    numbers[9] = "ninth"
    numbers[10] = "tenth"
    numbers[11] = "eleventh"
    numbers[12] = "twelfth"
    end := make(map[int]string)
    end[1] = "a Partridge in a Pear Tree"
    end[2] = "two Turtle Doves"
    end[3] = "three French Hens"
    end[4] = "four Calling Birds"
    end[5] = "five Gold Rings"
    end[6] = "six Geese-a-Laying"
    end[7] = "seven Swans-a-Swimming"
    end[8] = "eight Maids-a-Milking"
    end[9] = "nine Ladies Dancing"
    end[10] = "ten Lords-a-Leaping"
    end[11] = "eleven Pipers Piping"
    end[12] = "twelve Drummers Drumming"
    start := fmt.Sprintf("On the %s day of Christmas my true love gave to me: ", numbers[i])
    if i == 1 {
        return start + end[i] + "."
    }
    for j := i; j > 1; j-- {
        start += end[j] + ", "
    }
    return start + "and " + end[1] + "."
}

func Song() string {
    res := ""
	for i := 1; i < 12; i++ {
        res += Verse(i) + "\n"
    }
    res += Verse(12)
    return res
}
