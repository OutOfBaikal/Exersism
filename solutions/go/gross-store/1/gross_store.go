package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
    /*
    {"quarter_of_a_dozen", 3},
		{"half_of_a_dozen", 6},
		{"dozen", 12},
		{"small_gross", 120},
		{"gross", 144},
		{"great_gross", 1728}
    */
	res := make(map[string]int)
    res["quarter_of_a_dozen"] = 3
    res["half_of_a_dozen"] = 6
    res["dozen"] = 12
    res["small_gross"] = 120
    res["gross"] = 144
    res["great_gross"] = 1728
    return res
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int, 0)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	quantity, exists := units[unit]
    if !exists {
        return false
    }

    if _, found := bill[item]; found {
        bill[item] += quantity
    } else {
        bill[item] = quantity
    }

    return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, itemExists := bill[item]
	unitQuantity, unitExists := units[unit]
	if !itemExists || !unitExists {
		return false
	}

	bill[item] -= unitQuantity
    if bill[item] < 0 {
        bill[item] += unitQuantity
        return false
    } else if bill[item] == 0 {
        delete(bill, item)
    }
    return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	data, ok := bill[item]
    if !ok {
        return  0, false
    }
    return data, true
}

