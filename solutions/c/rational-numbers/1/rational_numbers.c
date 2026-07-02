#include "rational_numbers.h"
#include <stdlib.h>
#include <math.h>

static int32_t gcd(int32_t a, int32_t b) {
    a = abs(a);
    b = abs(b);
    while (b != 0) {
        int32_t temp = b;
        b = a % b;
        a = temp;
    }
    return a;
}

rational_t reduce(rational_t r) {
    int32_t common = gcd(r.numerator, r.denominator);
    r.numerator /= common;
    r.denominator /= common;

    if (r.denominator < 0) {
        r.numerator = -r.numerator;
        r.denominator = -r.denominator;
    }
    return r;
}

rational_t add(rational_t r1, rational_t r2) {
    rational_t result;
    result.numerator = r1.numerator * r2.denominator + r2.numerator * r1.denominator;
    result.denominator = r1.denominator * r2.denominator;
    return reduce(result);
}
rational_t subtract(rational_t r1, rational_t r2) {
    rational_t result;
    result.numerator = r1.numerator * r2.denominator - r2.numerator * r1.denominator;
    result.denominator = r1.denominator * r2.denominator;
    return reduce(result);
}
rational_t multiply(rational_t r1, rational_t r2) {
    rational_t result;
    result.numerator = r1.numerator * r2.numerator;
    result.denominator = r1.denominator * r2.denominator;
    return reduce(result);
}
rational_t divide(rational_t r1, rational_t r2) {
    rational_t result;
    result.numerator = r1.numerator * r2.denominator;
    result.denominator = r1.denominator * r2.numerator;
    return reduce(result);
}

rational_t absolute(rational_t r) {
    rational_t result = r;
    result.numerator = abs(result.numerator);
    result.denominator = abs(result.denominator);
    return reduce(result);
}

rational_t exp_rational(rational_t r, int16_t n) {
    rational_t result = {1, 1};

    if(n == 0) {
        return result;
    }
    r = reduce(r);
    if (n < 0) {
        rational_t inverted = {r.denominator, r.numerator};
        r = inverted;
        n = -n;
    }

    for (int16_t i = 0; i < n; ++i) {
        result = multiply(result, r);
    }

    return result;
}

float exp_real(int16_t x, rational_t r) {
    if (x <= 0) {
        return NAN; 
    }
    
    double exponent = (double)r.numerator / (double)r.denominator;
    return (float)pow((double)x, exponent);
}
