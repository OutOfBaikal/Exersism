package rationalnumbers

import "math"

type Rational struct {
	numerator, denominator int
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}
func gcd(a, b int) int {
    a = abs(a);
    b = abs(b);
    for b != 0 {
        a, b = b, a % b
    }
    return a
}
// Reduce simplifies a Rational, eg changing Rational{4, 2} into Rational{2, 1}.
func (r Rational) Reduce() Rational {
	common := gcd(r.numerator, r.denominator)
    r.numerator /= common
    r.denominator /= common
    if r.denominator < 0 {
        r.numerator = -r.numerator
        r.denominator = -r.denominator
    }
    return r
}

func (r Rational) Add(s Rational) Rational {
	var result Rational
    result.numerator = r.numerator * s.denominator + s.numerator * r.denominator
    result.denominator = r.denominator * s.denominator
    return result.Reduce()
}

func (r Rational) Sub(s Rational) Rational {
	var result Rational
    result.numerator = r.numerator * s.denominator - s.numerator * r.denominator
    result.denominator = r.denominator * s.denominator
    return result.Reduce()
}

func (r Rational) Mul(s Rational) Rational {
	var result Rational
    result.numerator = r.numerator * s.numerator
    result.denominator = r.denominator * s.denominator
    return result.Reduce()
}

func (r Rational) Div(s Rational) Rational {
	var result Rational
    result.numerator = r.numerator * s.denominator
    result.denominator = r.denominator * s.numerator
    return result.Reduce()
}

func (r Rational) Abs() Rational {
	result := r
    result.numerator = abs(result.numerator)
    result.denominator = abs(result.denominator)
    return result.Reduce()
}

// Compute r ^ power, a rational raised to an int exponent.
func (r Rational) Exprational(power int) Rational {
	var result Rational
    result.numerator = 1
    result.denominator = 1
    if power == 0 {
        return result
    }
    r = r.Reduce()
    if power < 0 {
        inverted := Rational{r.denominator, r.numerator}
        r = inverted
        power = -power
    }
    for i := 0; i < power; i++ {
        result = result.Mul(r)
    }
    return result
}

// Compute base ^ r, an int raised to a rational.
func (r Rational) Expreal(base int) float64 {
	if base <= 0 {
        return 0
    }
    exponent := float64(r.numerator) / float64(r.denominator);
    return math.Pow(float64(base), exponent)
}
