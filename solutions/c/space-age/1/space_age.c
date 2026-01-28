#include "space_age.h"
float age(planet_t planet, int64_t seconds) {
    float res = seconds / 31557600.0;
    switch(planet) {
        case 0:
            res /= 0.2408467;
            break;
        case 1:
            res /= 0.61519726;
            break;
        case 2:
            break;
        case 3:
            res /= 1.8808158;
            break;
        case 4:
            res /= 11.862615;
            break;
        case 5:
            res /= 29.447498;
            break;
        case 6:
            res /= 84.016846;
            break;
        case 7:
            res /= 164.79132;
            break;
        default:
            res = -1;
    }
    return res;
}