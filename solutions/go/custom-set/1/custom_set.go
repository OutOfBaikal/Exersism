package stringset

import (
    "fmt"
    "strings"
)

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Define the Set type here.

type Set map[string]struct{}

func New() Set {
	return make(Set)
}

func NewFromSlice(l []string) Set {
	stringSet := make(Set)
    for _, x := range l {
        stringSet[x] = struct{}{}
    }
    return stringSet
}

func (s Set) String() string {
	if len(s) == 0 {
        return "{}"
    }
    var elements []string
    for k := range s {
        elements = append(elements, fmt.Sprintf("%q", k))
    }
    return "{" + strings.Join(elements, ", ") + "}"
}

func (s Set) IsEmpty() bool {
	return len(s) == 0 
}

func (s Set) Has(elem string) bool {
	_, exist := s[elem]
    return exist
}

func (s Set) Add(elem string) {
	s[elem] = struct{}{}
}

func Subset(s1, s2 Set) bool {
	for k, _ := range s1 {
        if !s2.Has(k) {
			return false
		}
    }
    return true
}

func Disjoint(s1, s2 Set) bool {
    for k := range s1 {
		if s2.Has(k) {
			return false
		}
	}
	return true
}

func Equal(s1, s2 Set) bool {
	if len(s1) != len(s2) {
        return false
    }
    for k, _ := range s1 {
        if !s2.Has(k) {
			return false
		}
    }
    return true
}

func Intersection(s1, s2 Set) Set {
    result := New()
	for k := range s1 {
		if s2.Has(k) {
			result.Add(k)
		}
	}
    return result
}

func Difference(s1, s2 Set) Set {
    result := New()
	for k := range s1 {
		if !s2.Has(k) {
			result.Add(k)
		}
	}
    return result
}

func Union(s1, s2 Set) Set {
    result := New()
	for k := range s1 {
		result.Add(k)
	}
	for k := range s2 {
		result.Add(k)
	}
    return result
}
