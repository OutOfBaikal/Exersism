package sublist

// Relation type is defined in relations.go file.
func IsEqual(ListOne, ListTwo []int) bool {
    if len(ListOne) != len(ListTwo) {
        return false
    }
    for i := 0; i < len(ListOne); i++ {
        if ListOne[i] != ListTwo[i] {
            return false
        }
    }
    return true
}

func IsSuperlist(ListOne, ListTwo []int) bool {
    n, m := len(ListOne), len(ListTwo)
    for i := 0; i < n - m + 1; i++{
        if IsEqual(ListOne[i:i+m], ListTwo) {
            return true
        }
    }
    return false
}

func Sublist(l1, l2 []int) Relation {
    var result Relation
	if IsEqual(l1, l2) {
        result = "equal"
    } else if IsSuperlist(l1, l2) {
        result = "superlist"
    } else if IsSuperlist(l2, l1) {
        result = "sublist"
    } else {
        result = "unequal"
    }
    return result
}
