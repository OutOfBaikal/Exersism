package allergies

func Allergies(allergies uint) []string {
	data := []string{"eggs", "peanuts", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"}
    res := make([]string, 0)
    for i := 0; i < len(data); i++ {
        if ((allergies >> i) & 1) == 1 {
            res = append(res, data[i])
        }
    }
    return res
}

func AllergicTo(allergies uint, allergen string) bool {
	data := []string{"eggs", "peanuts", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"}
    var ind int
    for i, v := range(data) {
        if v == allergen {
            ind = i
            break
        }
    }
    return ((allergies >> ind) & 1) == 1
}
