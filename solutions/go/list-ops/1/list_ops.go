package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	for i := 0; i < len(s); i++ {
    	initial = fn(initial, s[i])
    }
    return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	for i := len(s) - 1; i >= 0; i-- {
		initial = fn(s[i], initial) 
	}
	return initial 
}

func (s IntList) Filter(fn func(int) bool) IntList {
	res := make(IntList, 0)
    for _, x := range(s) {
        if fn(x) {
            res = append(res, x)
        }
    }
    return res
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	res := make(IntList, 0)
    for _, x := range(s) {
        res = append(res, fn(x))
    }
    return res
}

func (s IntList) Reverse() IntList {
	for i := 0; i < len(s) / 2; i++ {
        s[i], s[len(s) - i - 1] = s[len(s) - i - 1], s[i]
    }
    return s
}

func (s IntList) Append(lst IntList) IntList {
	return append(s, lst...)
}

func (s IntList) Concat(lists []IntList) IntList {
	for _, x := range(lists) {
        s = append(s, x...)
    }
    return s
}
