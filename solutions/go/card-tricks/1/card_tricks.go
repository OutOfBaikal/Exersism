package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	res := make([]int, 3)
    res[0] = 2
    res[1] = 6
    res[2] = 9
    return res
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
    if index < len(slice) && index > -1 {
		return slice[index]
    } else {
        return -1
    }
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if index < len(slice) && index > -1 {
		slice[index] = value
    } else {
        slice = append(slice, value)
    }
    return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	res := make([]int, 0)
    res = append(res, values...)
    res = append(res, slice...)
    return res
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	res := make([]int, 0)
    if index > -1 && index < len(slice) {
    	res = append(slice[:index], slice[index+1:]...)
    } else {
        res = append(res, slice...)
    }
    return res
}
