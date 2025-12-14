#include "triangle.h"

bool is_equilateral(triangle_t sides) {
    double a = sides.a, b = sides.b, c = sides.c;
    if (a <= 0 || b <= 0 || c <= 0 || a + b <= c || a + c <= b || b + c <= a) {
        return false;
    }
    return (a == b && b == c);
}

bool is_isosceles(triangle_t sides) {
    double a = sides.a, b = sides.b, c = sides.c;
    if (a <= 0 || b <= 0 || c <= 0 || a + b <= c || a + c <= b || b + c <= a) {
        return false;
    }
    return (a == b || a == c || b == c);
}

bool is_scalene(triangle_t sides) {
    double a = sides.a, b = sides.b, c = sides.c;
    if (a <= 0 || b <= 0 || c <= 0 || a + b <= c || a + c <= b || b + c <= a) {
        return false;
    }
    return (a != b && a != c && b != c);
}