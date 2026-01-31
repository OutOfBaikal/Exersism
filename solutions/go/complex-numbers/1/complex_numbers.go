package complexnumbers

import (
    "math"
)

// Define the Number type here.
type Number struct {
    r float64
    i float64
}

func (n Number) Real() float64 {
	return n.r
}

func (n Number) Imaginary() float64 {
	return n.i
}

func (n1 Number) Add(n2 Number) Number {
	return Number{
        r: n1.r + n2.r,
        i: n1.i + n2.i,
    }
}

func (n1 Number) Subtract(n2 Number) Number {
    return Number{
        r: n1.r - n2.r,
        i: n1.i - n2.i,
    }
}

func (n1 Number) Multiply(n2 Number) Number {
	return Number{
        r: n1.r*n2.r - n1.i*n2.i,
        i: n1.i*n2.r + n1.r*n2.i,
    }
}

func (n Number) Times(factor float64) Number {
	return Number{
        r: n.r * factor,
        i: n.i * factor,
    }
}

func (n1 Number) Divide(n2 Number) Number {
	denom := n2.r*n2.r + n2.i*n2.i
    return Number{
        r: (n1.r*n2.r + n1.i*n2.i) / denom,
        i: (n1.i*n2.r - n1.r*n2.i) / denom,
    }
}

func (n Number) Conjugate() Number {
	return Number{
        r: n.r,
        i: n.i * -1,
    }
}

func (n Number) Abs() float64 {
	return math.Sqrt(n.r * n.r + n.i * n.i)
}

func (n Number) Exp() Number {
	return Number{
        r: math.Exp(n.r) * math.Cos(n.i),
        i: math.Exp(n.r) * math.Sin(n.i),
    }
}
