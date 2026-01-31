package isbn

func checkISBN(isbn string) bool {
    c := 0
    for i := 0; i < len(isbn); i++ {
        if isbn[i] >= '0' && isbn[i] <= '9' || isbn[i] == 'X' {
            c++
        }
    }
    if c < 10 {
        return false
    }
    if isbn[1] != '-' || isbn[5] != '-' || isbn[11] != '-' {
        return false
    }
    for i := 0; i < len(isbn); i++ {
        if i != 1 && i != 5 && i != 11 {
            if isbn[i] < '0' || isbn[i] > '9' {
                if i != 12 {
                    return false
                } else if isbn[i] != 'X' {
                    return false
                }
            }  
        }
    }
    return true
}

func IsValidISBN(isbn string) bool {
	if len(isbn) > 10 {
        if checkISBN(isbn) == false {
            return false
        }
    } else if len(isbn) < 10 {
        return false
    }
    sum := 0
    d := 10
    for i := 0; i < len(isbn); i++ {
        if isbn[i] >= 48 && isbn[i] <= 57 {
            sum += d * (int(isbn[i]) - 48)
            d--
        } else if isbn[i] == 'X' {
            sum += d * 10
            d--
        }
    }

    return sum % 11 == 0
}
