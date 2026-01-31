package sorting

import(
    "fmt"
    "strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

type myNumberBox struct {
    num int
}

func (n *myNumberBox) Number() int {
    return n.num
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// type differentFancyNumber struct {
// 	n string
// }

// func (i differentFancyNumber) Value() string {
// 	return i.n
// }

type differentFancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
    switch fnb.(type) {
        case FancyNumber:
        	var i, _ = strconv.Atoi(fnb.Value())
        	return i
        default:
        	return 0
    }
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	switch fnb.(type) {
        case FancyNumber:
        	var i, _ = strconv.Atoi(fnb.Value())
        	return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(i))
        default:
        	return fmt.Sprintf("This is a fancy box containing the number %.1f", 0.0)
    }
}

func DescribeDifferentFancyNumberBox(fnb differentFancyNumberBox) string {
	switch fnb.(type) {
        case differentFancyNumber:
        	var i, _ = strconv.Atoi(fnb.Value())
        	return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(i))
        default:
        	return fmt.Sprintf("This is a fancy box containing the number %.1f", 0.0)
    }
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	switch i.(type) {
		case int:
			return DescribeNumber(float64(i.(int)))
		case float64:
			return DescribeNumber(i.(float64))
		case NumberBox:
			return DescribeNumberBox(i.(NumberBox))
		case FancyNumber:
			return DescribeFancyNumberBox(i.(FancyNumber))
        case differentFancyNumber:
			return DescribeDifferentFancyNumberBox(i.(differentFancyNumber))
		default:
			return "Return to sender"
	}
}
