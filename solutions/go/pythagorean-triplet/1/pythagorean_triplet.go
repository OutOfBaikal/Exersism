package pythagorean

import(
    "math"
)

type Triplet [3]int

// Range generates list of all Pythagorean triplets with side lengths
// in the provided range.
func Range(min, max int) []Triplet {
	var res []Triplet
    for c := min; c <= max; c++ {
        for a := min; a <= c; a++ {
            b_sq := c * c - a * a;
            if b_sq <= 0 {
                continue
            }
            b := int(math.Sqrt(float64(b_sq)))
            if b >= a && b * b == b_sq && b <= max {
                res = append(res, Triplet{a, b, c})
            }
        }
    }
    return res
}

// Sum returns a list of all Pythagorean triplets with a certain perimeter.
func Sum(p int) []Triplet {
	var res []Triplet
    for a := 1; a < p; a++ {
        for b := a; b < p; b++ {
            c := p - a - b
            if c <= b {
                continue
            }
            if a * a + b * b == c * c {
                res = append(res, Triplet{a, b, c})
            }
        }
    }
    return res
}
