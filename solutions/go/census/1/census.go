// Package census simulates a system used to collect census data.
package census

func checkMapElements(m map[string]string) bool {
    if len(m) == 0 {
        return false
    }
    for k, v := range m {
        if k == "" || v == "" || k != "street" {
            return false
        }
    }
    return true
}

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
    return &Resident{
        Name: name,
        Age: age,
        Address: address,
    }
	panic("Please implement NewResident.")
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
    //something := make(map[string]string)
    if r.Name != "" && r.Address != nil && checkMapElements(r.Address) != false {
        return true
    }
    return false
	panic("Please implement HasRequiredInfo.")
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
    r.Name = ""
    r.Age = 0
    r.Address = nil
	//panic("Please implement Delete.")
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
    res := 0
    for i := 0; i < len(residents); i++ {
        if residents[i].HasRequiredInfo() == true {
            res++
        }
    }
    return res
	panic("Please implement Count.")
}
